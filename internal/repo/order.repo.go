package repo

import (
	"context"
	"database/sql"
	"math/rand"
	"time"

	"github.com/anle/codebase/internal/database"
	"github.com/anle/codebase/internal/model"
	uuidv4 "github.com/anle/codebase/internal/utils/uuid"
)

type IOrderRepo interface {
	GetOrders(ctx context.Context, input model.GetOrdersInput) (result []database.GetOrdersRow, err error)
	GetOrder(ctx context.Context, input model.GetOrderInput) (result database.GetOrderRow, err error)
	GetOrderItems(ctx context.Context, orderID string) (result []database.GetOrderItemsRow, err error)
	CreateOrder(ctx context.Context, input model.CreateOrderInput) (orderID string, err error)
	CreateOrderItem(ctx context.Context, input model.CreateOrderItemInput) (orderItemID string, err error)
	UpdateStatus(ctx context.Context, input model.UpdateStatusInput) (result sql.Result, err error)
	GetOrdersForAdmin(ctx context.Context, input model.GetOrdersForAdminInput) (result []database.GetOrdersForAdminRow, err error)
	GetOrderForAdmin(ctx context.Context, input model.GetOrderInput) (result database.GetOrderForAdminRow, err error)
	CreatePayment(ctx context.Context, input model.CreatePaymentInput) (paymentStatus bool, err error)
	GetOrderStatus(ctx context.Context, input model.GetOrderStatusInput) (result database.NullOrdersStatus, err error)
	GetOrderSummary(ctx context.Context) (result []database.GetOrderSummaryRow, err error)
	WithTx(tx *sql.Tx) IOrderRepo
}

type orderRepo struct {
	queries *database.Queries
}

// GetOrderSummary implements IOrderRepo.
func (or *orderRepo) GetOrderSummary(ctx context.Context) (result []database.GetOrderSummaryRow, err error) {
	result, err = or.queries.GetOrderSummary(ctx)
	if err != nil {
		return result, err
	}

	return result, nil
}

// GetOrderForAdmin implements IOrderRepo.
func (or *orderRepo) GetOrderForAdmin(ctx context.Context, input model.GetOrderInput) (result database.GetOrderForAdminRow, err error) {
	result, err = or.queries.GetOrderForAdmin(ctx, input.OrderID)
	if err != nil {
		return result, err
	}

	return result, nil
}

// GetOrderStatus implements IOrderRepo.
func (or *orderRepo) GetOrderStatus(ctx context.Context, input model.GetOrderStatusInput) (result database.NullOrdersStatus, err error) {
	result, err = or.queries.GetOrderStatus(ctx, database.GetOrderStatusParams{
		ID:     input.OrderID,
		UserID: ctx.Value("userID").(string),
	})

	if err != nil {
		return result, err
	}

	return result, nil
}

// CreatePayment implements IOrderRepo.
func (or *orderRepo) CreatePayment(ctx context.Context, input model.CreatePaymentInput) (paymentStatus bool, err error) {
	// Simulate third party payment
	// In a real-world scenario, you would call a payment gateway API here
	// and check the payment status.
	// For this example, we'll assume the payment is always successful.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// Update the order status to "Paid" in the database
	return r.Intn(2) == 1, nil
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
		Column5: input.OrderBy,
		Column6: input.OrderBy,
	})
	if err != nil {
		return result, err
	}

	return result, nil
}

// UpdateStatus implements IOrderRepo.
func (or *orderRepo) UpdateStatus(ctx context.Context, input model.UpdateStatusInput) (result sql.Result, err error) {
	result, err = or.queries.UpdateStatus(ctx, database.UpdateStatusParams{
		ID:     input.OrderID,
		Status: database.NullOrdersStatus{OrdersStatus: database.OrdersStatus(input.Status), Valid: input.Status != ""},
	})
	if err != nil {
		return result, err
	}

	return result, nil
}

// GetOrderItem implements IOrderRepo.
func (or *orderRepo) GetOrderItems(ctx context.Context, orderID string) (result []database.GetOrderItemsRow, err error) {
	result, err = or.queries.GetOrderItems(ctx, orderID)
	if err != nil {
		return result, err
	}

	return result, nil
}

// GetOrder implements IOrderRepo.
func (or *orderRepo) GetOrder(ctx context.Context, input model.GetOrderInput) (result database.GetOrderRow, err error) {
	result, err = or.queries.GetOrder(ctx, database.GetOrderParams{
		ID:     input.OrderID,
		UserID: ctx.Value("userID").(string),
	})
	if err != nil {
		return result, err
	}

	return result, nil
}

// GetOrders implements IOrderRepo.
func (or *orderRepo) GetOrders(ctx context.Context, input model.GetOrdersInput) (result []database.GetOrdersRow, err error) {
	result, err = or.queries.GetOrders(ctx, database.GetOrdersParams{
		UserID: ctx.Value("userID").(string),
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
func (or *orderRepo) CreateOrderItem(ctx context.Context, input model.CreateOrderItemInput) (orderItemID string, err error) {
	orderItemID = uuidv4.GenerateUUID()
	_, err = or.queries.CreateOrderItem(ctx, database.CreateOrderItemParams{
		ID:          orderItemID,
		OrderID:     input.OrderID,
		Name:        input.Name,
		Description: sql.NullString{String: input.Description, Valid: input.Description != ""},
		Price:       input.Price,
		Quantity:    input.Quantity,
		ImageUrl:    input.ImageUrl,
	})
	if err != nil {
		return orderItemID, err
	}

	return orderItemID, nil
}

// CreateOrder implements IOrderRepo.
func (or *orderRepo) CreateOrder(ctx context.Context, input model.CreateOrderInput) (orderID string, err error) {
	orderID = uuidv4.GenerateUUID()
	_, err = or.queries.CreateOrder(ctx, database.CreateOrderParams{
		ID:              orderID,
		UserID:          ctx.Value("userID").(string),
		PaymentMethod:   database.OrdersPaymentMethod(input.PaymentMethod),
		ShippingAddress: input.ShippingAddress,
		Total:           int64(input.TotalPrice),
	})
	if err != nil {
		return orderID, err
	}

	return orderID, nil
}

func NewOrderRepo(dbConn *sql.DB) IOrderRepo {
	return &orderRepo{
		queries: database.New(dbConn),
	}
}
