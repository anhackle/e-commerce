package repo

import (
	"context"
	"database/sql"

	"github.com/anle/codebase/internal/database"
	"github.com/anle/codebase/internal/model"
)

type IAuthenRepo interface {
	CreateUser(ctx context.Context, userInput model.UserInput) (err error)
	FindByEmail(ctx context.Context, input model.UserInput) (user database.User, err error)
}

type authenRepo struct {
	queries *database.Queries
}

func (ar *authenRepo) CreateUser(ctx context.Context, input model.UserInput) (err error) {
	_, err = ar.queries.CreateUser(ctx, database.CreateUserParams{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return err
	}

	return nil
}

func (ar *authenRepo) FindByEmail(ctx context.Context, input model.UserInput) (user database.User, err error) {
	user, err = ar.queries.FindByEmail(ctx, input.Email)
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
