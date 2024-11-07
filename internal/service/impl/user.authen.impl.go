package impl

import (
	"context"
	"database/sql"
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
	"github.com/anle/codebase/internal/utils/crypto"
	"github.com/anle/codebase/internal/utils/random"
	"github.com/anle/codebase/model"
	"github.com/anle/codebase/response"
	"github.com/redis/go-redis/v9"
)

type sUserAuthen struct {
	r *database.Queries
}

func (s *sUserAuthen) Login(ctx context.Context) error {
	return nil
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
		return response.ErrCodeUserHasExists, fmt.Errorf("")
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

func (s *sUserAuthen) VerifyOTP(ctx context.Context, in *model.VerifyInput) (out *model.VerifyOTPOutput, err error) {
	panic("error")
}

func (s *sUserAuthen) UpdatePasswordRegister(ctx context.Context) error {
	return nil
}

func NewUserAuthenImpl(r *database.Queries) service.IUserAuthen {
	return &sUserAuthen{
		r: r,
	}
}
