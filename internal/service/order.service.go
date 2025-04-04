package service

import (
	"context"
	"database/sql"
	"slices"
	"sort"

	"github.com/anle/codebase/internal/database"
	"github.com/anle/codebase/internal/model"
	"github.com/anle/codebase/internal/repo"
	"github.com/anle/codebase/response"
)

type IOrderService interface {
	GetOrders(ctx context.Context, input model.GetOrdersInput) (orders []model.GetOrdersOutput, result int, err error)
	CreateOrder(ctx context.Context, input model.CreateOrderInput) (result int, err error)
	GetOrder(ctx context.Context, input model.GetOrderInput) (orderDetail model.GetOrderOutput, result int, err error)
	UpdateStatus(ctx context.Context, input model.UpdateStatusInput) (result int, err error)
	GetOrdersForAdmin(ctx context.Context, input model.GetOrdersForAdminInput) (orders []model.GetOrdersForAdminOutput, result int, err error)
}

type orderService struct {
	db          *sql.DB
	cartRepo    repo.ICartRepo
	productRepo repo.IProductRepo
	orderRepo   repo.IOrderRepo
}

// GetOrdersForAdmin implements IOrderService.
func (os *orderService) GetOrdersForAdmin(ctx context.Context, input model.GetOrdersForAdminInput) (orders []model.GetOrdersForAdminOutput, result int, err error) {
	input.Page = (input.Page - 1) * input.Limit
	ordersRepo, err := os.orderRepo.GetOrdersForAdmin(ctx, input)
	if err != nil {
		return orders, response.ErrCodeInternal, err
	}

	for _, order := range ordersRepo {
		orders = append(orders, model.GetOrdersForAdminOutput{
			FirstName:        order.FirstName.String,
			LastName:         order.LastName.String,
			PhoneNumber:      order.PhoneNumber.String,
			CreatedAt:        order.CreatedAt.Time.Format("2006-01-02 15:04:05"),
			Status:           string(order.Status.OrdersStatus),
			ShippingAddreess: order.ShippingAddress,
			Payment_method:   string(order.PaymentMethod),
			Total:            order.Total,
		})
	}

	return orders, response.ErrCodeSuccess, nil
}

// UpdateStatus implements IOrderService.
func (os *orderService) UpdateStatus(ctx context.Context, input model.UpdateStatusInput) (result int, err error) {
	var status = map[string][]string{
		"pending":    {"paid", "cancelled"},
		"paid":       {"processing", "cancelled", "failed"},
		"processing": {"shipped"},
		"shipped":    {"delivered"},
		"delivered":  {},
		"cancelled":  {},
		"failed":     {},
	}
	order, err := os.orderRepo.GetOrder(ctx, model.GetOrderInput{
		OrderID: input.OrderID,
	})
	if err == sql.ErrNoRows {
		return response.ErrCodeOrderNotFound, err
	}

	if err != nil {
		return response.ErrCodeInternal, err
	}

	if verifyStatus := slices.Contains(status[string(order.Status.OrdersStatus)], input.Status); !verifyStatus {
		return response.ErrCodeStatusNotValid, err
	}

	_, err = os.orderRepo.UpdateStatus(ctx, input)
	if err != nil {
		return response.ErrCodeInternal, err
	}

	return response.ErrCodeSuccess, nil
}

// GetOrders implements IOrderService.
func (os *orderService) GetOrders(ctx context.Context, input model.GetOrdersInput) (orders []model.GetOrdersOutput, result int, err error) {
	var getOrdersInput = model.GetOrdersInput{
		Page:  (input.Page - 1) * input.Limit,
		Limit: input.Limit,
	}
	ordersRepo, err := os.orderRepo.GetOrders(ctx, getOrdersInput)
	if err != nil {
		return orders, response.ErrCodeInternal, err
	}

	for _, order := range ordersRepo {
		orders = append(orders, model.GetOrdersOutput{
			OrderID:          int(order.ID),
			CreatedAt:        order.CreatedAt.Time.Format("2006-01-02 15:04:05"),
			Status:           string(order.Status.OrdersStatus),
			ShippingAddreess: order.ShippingAddress,
			Payment_method:   string(order.PaymentMethod),
			Total:            order.Total,
		})
	}

	return orders, response.ErrCodeSuccess, nil
}

// CreateOrder implements IOrderService.
func (os *orderService) CreateOrder(ctx context.Context, input model.CreateOrderInput) (result int, err error) {
	//1. Get user's cart
	cart, err := os.cartRepo.GetCart(ctx)
	if err != nil {
		return response.ErrCodeInternal, err
	}

	if len(cart) == 0 {
		return response.ErrCodeCartEmpty, err
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

	//2.1 Sort the cart to prevent deadlock
	sort.Slice(cart, func(i, j int) bool {
		return cart[i].ProductID < cart[j].ProductID
	})

	//3. Check Stock
	var totalPrice int64 = 0
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

		totalPrice += int64(item.Quantity) * item.ProductPrice
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

// GetOrder implements IOrderService.
func (os *orderService) GetOrder(ctx context.Context, input model.GetOrderInput) (orderDetail model.GetOrderOutput, result int, err error) {
	order, err := os.orderRepo.GetOrder(ctx, input)
	if err == sql.ErrNoRows {
		return orderDetail, response.ErrCodeOrderNotFound, err
	}

	if err != nil {
		return orderDetail, response.ErrCodeInternal, err
	}

	items, err := os.orderRepo.GetOrderItems(ctx, int(order.ID))
	if err != nil {
		return orderDetail, response.ErrCodeInternal, err
	}

	for _, item := range items {
		orderDetail.OrderID = int(order.ID)
		orderDetail.CreatedAt = order.CreatedAt.Time.Format("2006-01-02 15:04:05")
		orderDetail.Status = string(order.Status.OrdersStatus)
		orderDetail.ShippingAddreess = order.ShippingAddress
		orderDetail.Payment_method = string(order.PaymentMethod)
		orderDetail.Total = order.Total
		orderDetail.Items = append(orderDetail.Items, model.GetOrderItemsOutput{
			Name:        item.Name,
			Description: item.Description.String,
			Price:       item.Price,
			Quantity:    int(item.Quantity),
			Image_url:   item.ImageUrl,
		})
	}

	return orderDetail, response.ErrCodeSuccess, nil
}

func NewOrderService(db *sql.DB, orderRepo repo.IOrderRepo, cartRepo repo.ICartRepo, productRepo repo.IProductRepo) IOrderService {
	return &orderService{
		db:          db,
		orderRepo:   orderRepo,
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}
