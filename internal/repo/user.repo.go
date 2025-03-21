package repo

import (
	"context"
	"database/sql"

	"github.com/anle/codebase/internal/database"
	"github.com/anle/codebase/internal/model"
)

type IUserRepo interface {
	UpdateProfile(ctx context.Context, input model.UpdateProfileInput) (err error)
	FindByUserId(ctx context.Context, userID int) (user database.User, err error)
	ChangePassword(ctx context.Context, newPassword string) (err error)
}

type userRepo struct {
	queries *database.Queries
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

func (ur *userRepo) FindByUserId(ctx context.Context, userID int) (user database.User, err error) {
	user, err = ur.queries.FindByUserId(ctx, int32(userID))
	if err != nil {
		return database.User{}, err
	}

	return user, nil
}

func NewUserRepo(dbConn *sql.DB) IUserRepo {
	return &userRepo{
		queries: database.New(dbConn),
	}
}
