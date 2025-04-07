package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/anle/codebase/internal/model"
	"github.com/anle/codebase/internal/repo"
	"github.com/anle/codebase/internal/utils/hash"
	"github.com/anle/codebase/response"
)

type IUserService interface {
	GetProfile(ctx context.Context) (profileResult model.GetProfileOutput, result int, err error)
	UpdateProfile(ctx context.Context, input model.UpdateProfileInput) (result int, err error)
	ChangePassword(ctx context.Context, input model.ChangePasswordInput) (result int, err error)
	UpdateRole(ctx context.Context, input model.UpdateRoleInput) (result int, err error)
	GetUsersForAdmin(ctx context.Context, input model.GetUsersForAdminInput) (users []model.GetUsersForAdminOutput, result int, err error)
	DeleteUser(ctx context.Context, input model.DeleteUserInput) (result int, err error)
}

type userService struct {
	userRepo repo.IUserRepo
}

// DeleteUser implements IUserService.
func (us *userService) DeleteUser(ctx context.Context, input model.DeleteUserInput) (result int, err error) {
	_, err = us.userRepo.DeleteUser(ctx, input)
	if err == sql.ErrNoRows {
		return response.ErrCodeUserNotFound, err
	}

	if err != nil {
		return response.ErrCodeInternal, err
	}

	return response.ErrCodeSuccess, nil
}

// GetUsersForAdmin implements IUserService.
func (us *userService) GetUsersForAdmin(ctx context.Context, input model.GetUsersForAdminInput) (users []model.GetUsersForAdminOutput, result int, err error) {
	input.Page = (input.Page - 1) * input.Limit
	usersRepo, err := us.userRepo.GetUsersForAdmin(ctx, input)
	if err != nil && err != sql.ErrNoRows {
		return users, response.ErrCodeInternal, err
	}

	for _, user := range usersRepo {
		users = append(users, model.GetUsersForAdminOutput{
			UserID: int(user.ID),
			Email:  user.Email,
			Role:   string(user.Role.UserRole),
		})
	}

	return users, response.ErrCodeSuccess, nil
}

// UpdateRole implements IUserService.
func (us *userService) UpdateRole(ctx context.Context, input model.UpdateRoleInput) (result int, err error) {
	_, err = us.userRepo.UpdateRole(ctx, input)
	if err == sql.ErrNoRows {
		return response.ErrCodeUserNotFound, err
	}

	if err != nil {
		return response.ErrCodeInternal, err
	}

	return response.ErrCodeSuccess, nil
}

// GetProfile implements IUserService.
func (us *userService) GetProfile(ctx context.Context) (profileResult model.GetProfileOutput, result int, err error) {
	user, err := us.userRepo.GetProfile(ctx, ctx.Value("userID").(int))
	if err != nil {
		return profileResult, response.ErrCodeInternal, err
	}

	profileResult = model.GetProfileOutput{
		FirstName:   user.FirstName.String,
		LastName:    user.LastName.String,
		PhoneNumber: user.PhoneNumber.String,
		Address:     user.Address.String,
	}

	return profileResult, response.ErrCodeSuccess, nil
}

func (us *userService) UpdateProfile(ctx context.Context, input model.UpdateProfileInput) (result int, err error) {
	err = us.userRepo.UpdateProfile(ctx, input)
	if err != nil {
		return response.ErrCodeInternal, err
	}

	return response.ErrCodeSuccess, nil
}

func (us *userService) ChangePassword(ctx context.Context, input model.ChangePasswordInput) (result int, err error) {
	user, err := us.userRepo.FindByUserId(ctx, ctx.Value("userID").(int))
	if err != nil && err != sql.ErrNoRows {
		return response.ErrCodeInternal, err
	}

	err = hash.ComparePassword(user.Password, input.OldPassword)
	if err != nil {
		return response.ErrCodeOldPasswordNotMatch, err
	}

	if input.NewPassword != input.ConfirmPassword {
		return response.ErrCodePasswordNotMatch, errors.New("new password and confirm password not match")
	}

	hashPassword, err := hash.HashPassword(input.NewPassword)
	if err != nil {
		return response.ErrCodeInternal, err
	}

	err = us.userRepo.ChangePassword(ctx, hashPassword)
	if err != nil {
		return response.ErrCodeInternal, err
	}

	return response.ErrCodeSuccess, nil
}

func NewUserService(userRepo repo.IUserRepo) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}
