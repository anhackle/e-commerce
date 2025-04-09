package repo

import (
	"context"
	"database/sql"

	"github.com/anle/codebase/internal/database"
	"github.com/anle/codebase/internal/model"
	uuidv4 "github.com/anle/codebase/internal/utils/uuid"
)

type ICartRepo interface {
	AddToCart(ctx context.Context, input model.AddToCartInput) (itemID string, err error)
	GetCart(ctx context.Context) (cart []database.GetCartRow, err error)
	DeleteCartByID(ctx context.Context, input model.DeleteCartInput) (result sql.Result, err error)
	DeleteCart(ctx context.Context) (result sql.Result, err error)
	UpdateCart(ctx context.Context, input model.UpdateCartInput) (result sql.Result, err error)
	WithTx(tx *sql.Tx) ICartRepo
}

type cartRepo struct {
	queries *database.Queries
}

// DeleteCart implements ICartRepo.
func (cr *cartRepo) DeleteCart(ctx context.Context) (result sql.Result, err error) {
	result, err = cr.queries.DeleteCart(ctx, ctx.Value("userID").(string))
	if err != nil {
		return result, err
	}

	return result, nil
}

// WithTx implements ICartRepo.
func (cr *cartRepo) WithTx(tx *sql.Tx) ICartRepo {
	return &cartRepo{
		queries: cr.queries.WithTx(tx),
	}
}

// UpdateCart implements ICartRepo.
func (cr *cartRepo) UpdateCart(ctx context.Context, input model.UpdateCartInput) (result sql.Result, err error) {
	result, err = cr.queries.UpdateCart(ctx, database.UpdateCartParams{
		ID:        input.ItemID,
		UserID:    ctx.Value("userID").(string),
		ProductID: input.ProductID,
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
func (cr *cartRepo) DeleteCartByID(ctx context.Context, input model.DeleteCartInput) (result sql.Result, err error) {
	result, err = cr.queries.DeleteCartByID(ctx, database.DeleteCartByIDParams{
		UserID: ctx.Value("userID").(string),
		ID:     input.ItemID,
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
func (cr *cartRepo) AddToCart(ctx context.Context, input model.AddToCartInput) (itemID string, err error) {
	itemID = uuidv4.GenerateUUID()
	_, err = cr.queries.AddToCart(ctx, database.AddToCartParams{
		ID:        itemID,
		UserID:    ctx.Value("userID").(string),
		ProductID: input.ProductID,
		Quantity:  int32(input.Quantity),
	})
	if err != nil {
		return itemID, err
	}

	return itemID, nil
}

// GetCart implements ICartRepo.
func (cr *cartRepo) GetCart(ctx context.Context) (cart []database.GetCartRow, err error) {
	cart, err = cr.queries.GetCart(ctx, ctx.Value("userID").(string))
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
