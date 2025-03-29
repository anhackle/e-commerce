package repo

import (
	"context"
	"database/sql"

	"github.com/anle/codebase/internal/database"
	"github.com/anle/codebase/internal/model"
)

type ICartRepo interface {
	AddToCart(ctx context.Context, input model.AddToCartInput) (result sql.Result, err error)
	GetCart(ctx context.Context) (cart []database.GetCartRow, err error)
	DeleteCart(ctx context.Context, input model.DeleteCartInput) (result sql.Result, err error)
	UpdateCart(ctx context.Context, input model.UpdateCartInput) (result sql.Result, err error)
}

type cartRepo struct {
	queries *database.Queries
}

// UpdateCart implements ICartRepo.
func (cr *cartRepo) UpdateCart(ctx context.Context, input model.UpdateCartInput) (result sql.Result, err error) {
	result, err = cr.queries.UpdateCart(ctx, database.UpdateCartParams{
		ID:        int32(input.ItemID),
		UserID:    int32(ctx.Value("userID").(int)),
		ProductID: int32(input.ProductID),
		Quantity:  int32(*input.Quantity),
	})
	if err != nil {
		return result, err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return result, sql.ErrNoRows
	}

	return result, nil

}

// DeleteCart implements ICartRepo.
func (cr *cartRepo) DeleteCart(ctx context.Context, input model.DeleteCartInput) (result sql.Result, err error) {
	result, err = cr.queries.DeleteCart(ctx, database.DeleteCartParams{
		UserID: int32(ctx.Value("userID").(int)),
		ID:     int32(input.ItemID),
	})
	if err != nil {
		return result, err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return result, sql.ErrNoRows
	}

	return result, nil
}

// AddToCart implements ICartRepo.
func (cr *cartRepo) AddToCart(ctx context.Context, input model.AddToCartInput) (result sql.Result, err error) {
	result, err = cr.queries.AddToCart(ctx, database.AddToCartParams{
		UserID:    int32(ctx.Value("userID").(int)),
		ProductID: int32(input.ProductID),
		Quantity:  int32(input.Quantity),
	})
	if err != nil {
		return result, err
	}

	return result, nil
}

// GetCart implements ICartRepo.
func (cr *cartRepo) GetCart(ctx context.Context) (cart []database.GetCartRow, err error) {
	cart, err = cr.queries.GetCart(ctx, int32(ctx.Value("userID").(int)))
	if err != nil {
		return cart, err
	}

	return cart, nil
}

func NewCartRepo(dbConn *sql.DB) ICartRepo {
	return &cartRepo{
		queries: database.New(dbConn),
	}
}
