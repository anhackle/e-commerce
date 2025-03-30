package repo

import (
	"context"
	"database/sql"

	"github.com/anle/codebase/internal/database"
	"github.com/anle/codebase/internal/model"
)

type IOrderRepo interface {
	GetOrders(ctx context.Context, input model.GetOrdersInput) (result []database.GetOrdersRow, err error)
	GetOrder(ctx context.Context, input model.GetOrderInput) (result database.GetOrderRow, err error)
	GetOrderItems(ctx context.Context, orderID int) (result []database.GetOrderItemsRow, err error)
	CreateOrder(ctx context.Context, input model.CreateOrderInput) (result sql.Result, err error)
	CreateOrderItem(ctx context.Context, input model.CreateOrderItemInput) (result sql.Result, err error)
	UpdateStatus(ctx context.Context, input model.UpdateStatusInput) (result sql.Result, err error)
	GetOrdersForAdmin(ctx context.Context, input model.GetOrdersForAdminInput) (result []database.GetOrdersForAdminRow, err error)
	WithTx(tx *sql.Tx) IOrderRepo
}

type orderRepo struct {
	queries *database.Queries
}

// GetOrdersForAdmin implements IOrderRepo.
func (or *orderRepo) GetOrdersForAdmin(ctx context.Context, input model.GetOrdersForAdminInput) (result []database.GetOrdersForAdminRow, err error) {
	result, err = or.queries.GetOrdersForAdmin(ctx, database.GetOrdersForAdminParams{
		Limit:   int32(input.Limit),
		Offset:  int32(input.Page),
		Column1: input.Status,
		IF:      input.Status,
		Column3: input.Payment_method,
		IF_2:    input.Payment_method,
	})
	if err != nil {
		return result, err
	}

	return result, nil
}

// UpdateStatus implements IOrderRepo.
func (or *orderRepo) UpdateStatus(ctx context.Context, input model.UpdateStatusInput) (result sql.Result, err error) {
	result, err = or.queries.UpdateStatus(ctx, database.UpdateStatusParams{
		ID:     int32(input.OrderID),
		Status: database.NullOrdersStatus{OrdersStatus: database.OrdersStatus(input.Status), Valid: input.Status != ""},
	})
	if err != nil {
		return result, err
	}

	return result, nil
}

// GetOrderItem implements IOrderRepo.
func (or *orderRepo) GetOrderItems(ctx context.Context, orderID int) (result []database.GetOrderItemsRow, err error) {
	result, err = or.queries.GetOrderItems(ctx, int32(orderID))
	if err != nil {
		return result, err
	}

	return result, nil
}

// GetOrder implements IOrderRepo.
func (or *orderRepo) GetOrder(ctx context.Context, input model.GetOrderInput) (result database.GetOrderRow, err error) {
	result, err = or.queries.GetOrder(ctx, database.GetOrderParams{
		ID:     int32(input.OrderID),
		UserID: int32(ctx.Value("userID").(int)),
	})
	if err != nil {
		return result, err
	}

	return result, nil
}

// GetOrders implements IOrderRepo.
func (or *orderRepo) GetOrders(ctx context.Context, input model.GetOrdersInput) (result []database.GetOrdersRow, err error) {
	result, err = or.queries.GetOrders(ctx, database.GetOrdersParams{
		UserID: int32(ctx.Value("userID").(int)),
		Limit:  int32(input.Limit),
		Offset: int32(input.Page),
	})
	if err != nil {
		return result, err
	}

	return result, nil
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
		PaymentMethod:   database.OrdersPaymentMethod(input.PaymentMethod),
		ShippingAddress: input.ShippingAddress,
		Total:           int64(input.TotalPrice),
	})
	if err != nil {
		return result, err
	}

	return result, nil
}

func NewOrderRepo(dbConn *sql.DB) IOrderRepo {
	return &orderRepo{
		queries: database.New(dbConn),
	}
}
