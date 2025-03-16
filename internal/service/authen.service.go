package service

import (
	"context"
	"database/sql"

	"github.com/anle/codebase/internal/model"
	"github.com/anle/codebase/internal/repo"
	"github.com/anle/codebase/internal/utils/hash"
	"github.com/anle/codebase/response"
)

type IAuthenService interface {
	Register(ctx context.Context, input model.UserInput) (result int, err error)
}

type authenService struct {
	authenRepo repo.IAuthenRepo
}

func (as *authenService) Register(ctx context.Context, input model.UserInput) (result int, err error) {
	_, err = as.authenRepo.FindByEmail(ctx, input)
	if err == nil {
		return response.ErrCodeUserHasExists, nil
	}

	if err != nil && err != sql.ErrNoRows {
		return response.ErrCodeInternal, err
	}

	hashPassword, err := hash.Hash(input.Password)
	if err != nil {
		return response.ErrCodeInternal, err
	}

	var userInput = model.UserInput{
		Email:    input.Email,
		Password: hashPassword,
	}
	err = as.authenRepo.CreateUser(ctx, userInput)
	if err != nil {
		return response.ErrCodeInternal, err
	}

	return response.ErrCodeSuccess, nil
}

func NewAuthenService(authenRepo repo.IAuthenRepo) IAuthenService {
	return &authenService{
		authenRepo: authenRepo,
	}
}
