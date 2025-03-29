package repo

import (
	"context"
	"database/sql"

	"github.com/anle/codebase/internal/database"
	"github.com/anle/codebase/internal/model"
)

type IOrderRepo interface {
	GetOrder(ctx context.Context) (err error)
	CreateOrder(ctx context.Context, input model.CreateOrderInput) (result sql.Result, err error)
	CreateOrderItem(ctx context.Context, input model.CreateOrderItemInput) (result sql.Result, err error)
	DeleteOrder(ctx context.Context) (err error)
	WithTx(tx *sql.Tx) IOrderRepo
}

type orderRepo struct {
	queries *database.Queries
}

// WithTx implements IOrderRepo.
func (or *orderRepo) WithTx(tx *sql.Tx) IOrderRepo {
	return &orderRepo{
		queries: or.queries.WithTx(tx),
	}
}

// CreateOrderItem implements IOrderRepo.
func (or *orderRepo) CreateOrderItem(ctx context.Context, input model.CreateOrderItemInput) (result sql.Result, err error) {
	result, err = or.queries.CreateOrderItem(ctx, database.CreateOrderItemParams{
		OrderID:     input.OrderID,
		Name:        input.Name,
		Description: sql.NullString{String: input.Description, Valid: input.Description != ""},
		Price:       input.Price,
		Quantity:    input.Quantity,
		ImageUrl:    input.ImageUrl,
	})
	if err != nil {
		return result, err
	}

	return result, nil
}

// CreateOrder implements IOrderRepo.
func (or *orderRepo) CreateOrder(ctx context.Context, input model.CreateOrderInput) (result sql.Result, err error) {
	result, err = or.queries.CreateOrder(ctx, database.CreateOrderParams{
		UserID:          int32(ctx.Value("userID").(int)),
		PaymentMethod:   database.OrderPaymentMethod(input.PaymentMethod),
		ShippingAddress: input.ShippingAddress,
		Total:           int64(input.TotalPrice),
	})
	if err != nil {
		return result, err
	}

	return result, nil
}

// DeleteOrder implements IOrderRepo.
func (or *orderRepo) DeleteOrder(ctx context.Context) (err error) {
	panic("unimplemented")
}

// GetOrder implements IOrderRepo.
func (or *orderRepo) GetOrder(ctx context.Context) (err error) {
	panic("unimplemented")
}

func NewOrderRepo(dbConn *sql.DB) IOrderRepo {
	return &orderRepo{
		queries: database.New(dbConn),
	}
}
