package impl

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/consts"
	"github.com/anle/codebase/internal/database"
	"github.com/anle/codebase/internal/service"
	"github.com/anle/codebase/internal/utils"
	"github.com/anle/codebase/internal/utils/auth"
	"github.com/anle/codebase/internal/utils/crypto"
	"github.com/anle/codebase/internal/utils/random"
	"github.com/anle/codebase/model"
	"github.com/anle/codebase/response"
	"github.com/redis/go-redis/v9"
)

type sUserAuthen struct {
	r *database.Queries
}

func (s *sUserAuthen) Login(ctx context.Context, in *model.LoginInput) (codeResult int, out model.LoginOutput, err error) {
	userBase, err := s.r.GetOneUserInfoAdmin(ctx, in.UserAccount)
	if err != nil {
		return response.ErrCodeAccountNotExist, model.LoginOutput{}, err
	}

	if !crypto.MatchingPassword(userBase.UserPassword, in.UserPassword, userBase.UserSalt) {
		return response.ErrCodeAccountNotExist, out, fmt.Errorf("does not match password")
	}

	//TODO: Check two-factor authentication

	go s.r.LoginUserBase(ctx, database.LoginUserBaseParams{
		UserLoginIp:  sql.NullString{String: "127.0.0.1", Valid: true},
		UserAccount:  in.UserAccount,
		UserPassword: in.UserPassword,
	})

	subToken := utils.GenerateClientTokenUUID(int(userBase.UserID))
	log.Println("subtoken:", subToken)
	infoUser, err := s.r.GetUser(ctx, uint64(userBase.UserID))
	if err != nil {
		return response.ErrCodeAccountNotExist, out, err
	}

	infoUserJson, err := json.Marshal(infoUser)
	if err != nil {
		return response.ErrCodeAccountNotExist, out, fmt.Errorf("convert to json fail")
	}

	err = global.Rdb.Set(ctx, subToken, infoUserJson, time.Duration(consts.TIME_OTP_REGISTER)*time.Minute).Err()
	if err != nil {
		return response.ErrCodeAccountNotExist, out, err
	}

	out.Token, err = auth.CreateToken(subToken)
	if err != nil {
		return response.ErrCodeAccountNotExist, out, err
	}

	return 200, out, nil

}

func (s *sUserAuthen) Register(ctx context.Context, in *model.RegisterInput) (codeResult int, err error) {
	//TODO
	//1. Hash email
	fmt.Printf("VerifyKey: %s\n", in.VerifyKey)
	fmt.Printf("VerifyKey:%d\n", in.VerifyType)
	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))
	fmt.Printf("hashKey: %s\n", hashKey)

	//2. Check user exists in user base
	userFound, err := s.r.CheckUserBaseExists(ctx, in.VerifyKey)
	if err != nil {
		return response.ErrCodeUserHasExists, err
	}

	if userFound > 0 {
		return response.ErrCodeUserHasExists, fmt.Errorf("user has already registered")
	}

	//3. Create OTP
	userKey := utils.GetUserKey(hashKey)
	otpFound, err := global.Rdb.Get(ctx, userKey).Result()

	switch {
	case err == redis.Nil:
		fmt.Println("Key does not exist")
	case err != nil:
		fmt.Println("get failed::", err)
	case otpFound != "":
		return response.ErrCodeUserHasExists, fmt.Errorf("please get otp after 1 minute")
	}

	//4. Generate OTP
	otpNew := random.GenerateSixDigitOTP()
	if in.VerifyPurpose == "TEST_USER" {
		otpNew = 123456
	}
	fmt.Printf("otp is :::%d\n", otpNew)

	//5. Save OTP in Redis with expiration time
	err = global.Rdb.SetEx(ctx, userKey, strconv.Itoa(otpNew), time.Duration(consts.TIME_OTP_REGISTER)*time.Minute).Err()
	if err != nil {
		return response.ErrCodeOTPInvalid, err
	}

	//6. Send OTP
	switch in.VerifyType {
	case consts.EMAIL:
		fmt.Println("Send OTP to email successfully")
		//7. save OTP to Mysql:
		result, err := s.r.InsertOTPVerify(ctx, database.InsertOTPVerifyParams{
			VerifyOtp:     strconv.Itoa(otpNew),
			VerifyType:    sql.NullInt32{Int32: 1, Valid: true},
			VerifyKey:     in.VerifyKey,
			VerifyKeyHash: hashKey,
		})

		if err != nil {
			return response.ErrCodeOTPInvalid, err
		}

		//8. getlastID
		lastIdVerifyUser, err := result.LastInsertId()
		if err != nil {
			return response.ErrCodeOTPInvalid, err
		}
		log.Println("lastIdVerifyUser", lastIdVerifyUser)
		return response.ErrCodeSuccess, nil

	case consts.MOBILE:
		fmt.Printf("Send OTP through SMS")
	}

	return response.ErrCodeSuccess, nil
}

func (s *sUserAuthen) VerifyOTP(ctx context.Context, in *model.VerifyInput) (out model.VerifyOTPOutput, err error) {
	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))

	otpFound, err := global.Rdb.Get(ctx, utils.GetUserKey(hashKey)).Result()
	if err != nil {
		return out, err
	}

	if in.VerifyCode != otpFound {
		// TODO: 3 time wrong OTP in 1 minute
		return out, fmt.Errorf("OTP not match")
	}

	infoOTP, err := s.r.GetInfoOTP(ctx, hashKey)
	if err != nil {
		return out, err
	}

	err = s.r.UpdateUserVerifiationStatus(ctx, hashKey)
	if err != nil {
		return out, err
	}

	// fmt.Println(out.Message)
	// out = &model.VerifyOTPOutput{
	// 	Token:   infoOTP.VerifyKeyHash,
	// 	Message: "success",
	// }

	out.Token = infoOTP.VerifyKeyHash
	out.Message = "success"

	return out, err
}

func (s *sUserAuthen) UpdatePasswordRegister(ctx context.Context, token, password string) (userId int, err error) {
	infoOTP, err := s.r.GetInfoOTP(ctx, token)
	if err != nil {
		return userId, fmt.Errorf("OTP not exists")
	}

	if infoOTP.IsVerified.Int32 == 0 {
		return userId, fmt.Errorf("user OTP not verified")
	}

	userBase := database.AddUserBaseParams{}
	userBase.UserAccount = infoOTP.VerifyKey
	userSalt, err := crypto.GenerateSalt(16)
	if err != nil {
		return response.ErrCodeOTPInvalid, nil
	}

	userBase.UserSalt = userSalt
	userBase.UserPassword = crypto.HashPassword(password, userSalt)
	newUserBase, err := s.r.AddUserBase(ctx, userBase)
	if err != nil {
		return response.ErrCodeOTPInvalid, err
	}
	user_id, err := newUserBase.LastInsertId()
	if err != nil {
		return response.ErrCodeOTPInvalid, err
	}

	return int(user_id), nil

}

func NewUserAuthenImpl(r *database.Queries) service.IUserAuthen {
	return &sUserAuthen{
		r: r,
	}
}
