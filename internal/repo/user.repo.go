package repo

import (
	"context"
	"database/sql"

	"github.com/anle/codebase/internal/database"
	"github.com/anle/codebase/internal/model"
)

type IUserRepo interface {
	UpdateProfile(ctx context.Context, input model.UpdateProfileInput) (err error)
	GetProfile(ctx context.Context, userId int) (user database.GetUserProfileRow, err error)
	FindByUserId(ctx context.Context, userID int) (user database.FindByUserIdRow, err error)
	ChangePassword(ctx context.Context, newPassword string) (err error)
	UpdateRole(ctx context.Context, input model.UpdateRoleInput) (result sql.Result, err error)
}

type userRepo struct {
	queries *database.Queries
}

// UpdateRole implements IUserRepo.
func (ur *userRepo) UpdateRole(ctx context.Context, input model.UpdateRoleInput) (result sql.Result, err error) {
	result, err = ur.queries.UpdateRole(ctx, database.UpdateRoleParams{
		ID:   int32(input.UserID),
		Role: database.NullUserRole{UserRole: database.UserRole(input.Role), Valid: input.Role != ""},
	})
	if err != nil {
		return result, err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return result, sql.ErrNoRows
	}

	return result, nil
}

// GetProfile implements IUserRepo.
func (ur *userRepo) GetProfile(ctx context.Context, userId int) (user database.GetUserProfileRow, err error) {
	user, err = ur.queries.GetUserProfile(ctx, int32(userId))
	if err != nil {
		return database.GetUserProfileRow{}, err
	}

	return user, nil
}

func (ur *userRepo) ChangePassword(ctx context.Context, newPassword string) (err error) {
	_, err = ur.queries.UpdatePassword(ctx, database.UpdatePasswordParams{
		Password: newPassword,
		ID:       int32(ctx.Value("userID").(int)),
	})
	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepo) UpdateProfile(ctx context.Context, input model.UpdateProfileInput) (err error) {
	_, err = ur.queries.UpdateUserProfile(ctx, database.UpdateUserProfileParams{
		FirstName:   sql.NullString{String: input.FirstName, Valid: input.FirstName != ""},
		LastName:    sql.NullString{String: input.LastName, Valid: input.LastName != ""},
		PhoneNumber: sql.NullString{String: input.PhoneNumber, Valid: input.PhoneNumber != ""},
		Address:     sql.NullString{String: input.Address, Valid: input.Address != ""},
		UserID:      int32(ctx.Value("userID").(int)),
	})
	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepo) FindByUserId(ctx context.Context, userID int) (user database.FindByUserIdRow, err error) {
	user, err = ur.queries.FindByUserId(ctx, int32(userID))
	if err != nil {
		return database.FindByUserIdRow{}, err
	}

	return user, nil
}

func NewUserRepo(dbConn *sql.DB) IUserRepo {
	return &userRepo{
		queries: database.New(dbConn),
	}
}
