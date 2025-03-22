package repo

import (
	"context"
	"database/sql"

	"github.com/anle/codebase/internal/database"
	"github.com/anle/codebase/internal/model"
)

type IAuthenRepo interface {
	CreateUser(ctx context.Context, input model.RegisterInput) (result sql.Result, err error)
	CreateUserProfile(ctx context.Context, userID int) (err error)
	FindByEmail(ctx context.Context, input string) (user database.User, err error)
	WithTx(tx *sql.Tx) IAuthenRepo
}

type authenRepo struct {
	queries *database.Queries
}

func (ar *authenRepo) WithTx(tx *sql.Tx) IAuthenRepo {
	return &authenRepo{
		queries: ar.queries.WithTx(tx),
	}
}

// CreateUserProfile implements IAuthenRepo.
func (ar *authenRepo) CreateUserProfile(ctx context.Context, userID int) (err error) {
	_, err = ar.queries.CreateUserProfile(ctx, int32(userID))
	if err != nil {
		return err
	}

	return nil
}

func (ar *authenRepo) CreateUser(ctx context.Context, input model.RegisterInput) (result sql.Result, err error) {
	result, err = ar.queries.CreateUser(ctx, database.CreateUserParams{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return result, err
	}

	return result, nil
}

func (ar *authenRepo) FindByEmail(ctx context.Context, input string) (user database.User, err error) {
	user, err = ar.queries.FindByEmail(ctx, input)
	if err != nil {
		return database.User{}, err
	}

	return user, nil
}

func NewAuthenRepo(dbConn *sql.DB) IAuthenRepo {
	return &authenRepo{
		queries: database.New(dbConn),
	}
}
