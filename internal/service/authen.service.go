package service

import (
	"context"
	"database/sql"

	"github.com/anle/codebase/internal/model"
	"github.com/anle/codebase/internal/repo"
	"github.com/anle/codebase/internal/utils/hash"
	"github.com/anle/codebase/internal/utils/jwttoken"
	"github.com/anle/codebase/response"
)

type IAuthenService interface {
	Register(ctx context.Context, input model.RegisterInput) (result int, err error)
	Login(ctx context.Context, input model.LoginInput) (token string, result int, err error)
}

type authenService struct {
	db         *sql.DB
	authenRepo repo.IAuthenRepo
}

func (as *authenService) Register(ctx context.Context, input model.RegisterInput) (result int, err error) {
	_, err = as.authenRepo.FindByEmail(ctx, input.Email)
	if err == nil {
		return response.ErrCodeUserHasExists, nil
	}

	if err != nil && err != sql.ErrNoRows {
		return response.ErrCodeInternal, err
	}

	hashPassword, err := hash.HashPassword(input.Password)
	if err != nil {
		return response.ErrCodeInternal, err
	}

	tx, err := as.db.Begin()
	if err != nil {
		return response.ErrCodeInternal, err
	}
	txRepo := as.authenRepo.WithTx(tx)
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		}
	}()

	var userInput = model.RegisterInput{
		Email:    input.Email,
		Password: hashPassword,
	}
	userID, err := txRepo.CreateUser(ctx, userInput)
	if err != nil {
		tx.Rollback()
		return response.ErrCodeInternal, err
	}

	if err != nil {
		tx.Rollback()
		return response.ErrCodeInternal, err
	}

	err = txRepo.CreateUserProfile(ctx, userID)
	if err != nil {
		tx.Rollback()
		return response.ErrCodeInternal, err
	}

	if err := tx.Commit(); err != nil {
		return response.ErrCodeInternal, err
	}
	return response.ErrCodeSuccess, nil
}

func (as *authenService) Login(ctx context.Context, input model.LoginInput) (token string, result int, err error) {
	user, err := as.authenRepo.FindByEmail(ctx, input.Email)
	if err != nil && err != sql.ErrNoRows {
		return "", response.ErrCodeInternal, err
	}

	err = hash.ComparePassword(user.Password, input.Password)
	if err != nil {
		return "", response.ErrCodeNotAuthorize, nil
	}

	token, err = jwttoken.GenJWTToken(user.ID, string(user.Role.UserRole))
	if err != nil {
		return "", response.ErrCodeInternal, err
	}

	return token, response.ErrCodeSuccess, nil
}

func NewAuthenService(db *sql.DB, authenRepo repo.IAuthenRepo) IAuthenService {
	return &authenService{
		db:         db,
		authenRepo: authenRepo,
	}
}
