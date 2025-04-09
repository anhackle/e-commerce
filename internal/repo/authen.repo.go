package repo

import (
	"context"
	"database/sql"

	"github.com/anle/codebase/internal/database"
	"github.com/anle/codebase/internal/model"
	uuidv4 "github.com/anle/codebase/internal/utils/uuid"
)

type IAuthenRepo interface {
	CreateUser(ctx context.Context, input model.RegisterInput) (userID string, err error)
	CreateUserProfile(ctx context.Context, userID string) (err error)
	FindByEmail(ctx context.Context, input string) (user database.FindByEmailRow, err error)
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
func (ar *authenRepo) CreateUserProfile(ctx context.Context, userID string) (err error) {
	_, err = ar.queries.CreateUserProfile(ctx, database.CreateUserProfileParams{
		ID:     uuidv4.GenerateUUID(),
		UserID: userID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (ar *authenRepo) CreateUser(ctx context.Context, input model.RegisterInput) (userID string, err error) {
	userID = uuidv4.GenerateUUID()
	_, err = ar.queries.CreateUser(ctx, database.CreateUserParams{
		ID:       userID,
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return userID, err
	}

	return userID, nil
}

func (ar *authenRepo) FindByEmail(ctx context.Context, input string) (user database.FindByEmailRow, err error) {
	user, err = ar.queries.FindByEmail(ctx, input)
	if err != nil {
		return database.FindByEmailRow{}, err
	}

	return user, nil
}

func NewAuthenRepo(dbConn *sql.DB) IAuthenRepo {
	return &authenRepo{
		queries: database.New(dbConn),
	}
}
