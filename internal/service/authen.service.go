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

	var userInput = model.RegisterInput{
		Email:    input.Email,
		Password: hashPassword,
	}
	sqlResult, err := as.authenRepo.CreateUser(ctx, userInput)
	if err != nil {
		return response.ErrCodeInternal, err
	}

	userID, err := sqlResult.LastInsertId()
	if err != nil {
		return response.ErrCodeInternal, err
	}

	err = as.authenRepo.CreateUserProfile(ctx, int(userID))
	if err != nil {
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

	token, err = jwttoken.GenJWTToken(int(user.ID))
	if err != nil {
		return "", response.ErrCodeInternal, err
	}

	return token, response.ErrCodeSuccess, nil
}

func NewAuthenService(authenRepo repo.IAuthenRepo) IAuthenService {
	return &authenService{
		authenRepo: authenRepo,
	}
}
