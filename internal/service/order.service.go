package service

import (
	"context"
	"database/sql"

	"github.com/anle/codebase/internal/database"
	"github.com/anle/codebase/internal/model"
	"github.com/anle/codebase/internal/repo"
	"github.com/anle/codebase/response"
)

type IOrderService interface {
	GetOrder(ctx context.Context) (result int, err error)
	CreateOrder(ctx context.Context, input model.CreateOrderInput) (result int, err error)
	DeleteOrder(ctx context.Context) (result int, err error)
}

type orderService struct {
	db          *sql.DB
	cartRepo    repo.ICartRepo
	productRepo repo.IProductRepo
	orderRepo   repo.IOrderRepo
}

// CreateOrder implements IOrderService.
func (os *orderService) CreateOrder(ctx context.Context, input model.CreateOrderInput) (result int, err error) {
	//1. Get user's cart
	cart, err := os.cartRepo.GetCart(ctx)
	if err == sql.ErrNoRows {
		return response.ErrCodeCartEmpty, err
	}

	if err != nil {
		return response.ErrCodeInternal, err
	}

	//2. Create transaction
	tx, err := os.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: 4,
	})
	if err != nil {
		return response.ErrCodeInternal, err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	//3. Check Stock
	totalPrice := 0
	products := make(map[int32]database.GetProductByIDForUpdateRow)

	for _, item := range cart {
		product, err := os.productRepo.WithTx(tx).GetProductByIDForUpdate(ctx, int(item.ProductID))
		if err == sql.ErrNoRows {
			tx.Rollback()
			return response.ErrCodeExternal, err
		}

		if err != nil {
			tx.Rollback()
			return response.ErrCodeInternal, err
		}

		if product.Quantity < item.Quantity {
			tx.Rollback()
			return response.ErrCodeQuantityNotEnough, err
		}

		totalPrice += int(item.Total)
		products[product.ID] = product
	}

	//4. Create order
	createOrderResult, err := os.orderRepo.WithTx(tx).CreateOrder(ctx, model.CreateOrderInput{
		PaymentMethod:   input.PaymentMethod,
		ShippingAddress: input.ShippingAddress,
		TotalPrice:      totalPrice,
	})
	if err != nil {
		tx.Rollback()
		return response.ErrCodeInternal, err
	}

	//5. Create order_items
	orderID, err := createOrderResult.LastInsertId()
	if err != nil {
		tx.Rollback()
		return response.ErrCodeInternal, err
	}

	for _, item := range cart {
		_, err = os.orderRepo.WithTx(tx).CreateOrderItem(ctx, model.CreateOrderItemInput{
			OrderID:     int32(orderID),
			Name:        products[item.ProductID].Name,
			Description: products[item.ProductID].Description.String,
			Price:       products[item.ProductID].Price,
			Quantity:    item.Quantity,
			ImageUrl:    products[item.ProductID].ImageUrl,
		})
		if err != nil {
			tx.Rollback()
			return response.ErrCodeInternal, err
		}

		_, err = os.productRepo.WithTx(tx).UpdateProductByID(ctx, model.UpdateProductByIDInput{
			ID:       int(item.ProductID),
			Quantity: int(products[item.ProductID].Quantity) - int(item.Quantity),
		})
		if err != nil {
			tx.Rollback()
			return response.ErrCodeInternal, err
		}
	}

	//6. Delete cart
	_, err = os.cartRepo.WithTx(tx).DeleteCart(ctx)
	if err != nil {
		tx.Rollback()
		return response.ErrCodeExternal, err
	}

	tx.Commit()

	return response.ErrCodeSuccess, err
}

// DeleteOrder implements IOrderService.
func (os *orderService) DeleteOrder(ctx context.Context) (result int, err error) {
	panic("unimplemented")
}

// GetOrder implements IOrderService.
func (os *orderService) GetOrder(ctx context.Context) (result int, err error) {
	panic("unimplemented")
}

func NewOrderService(db *sql.DB, orderRepo repo.IOrderRepo, cartRepo repo.ICartRepo, productRepo repo.IProductRepo) IOrderService {
	return &orderService{
		db:          db,
		orderRepo:   orderRepo,
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}
