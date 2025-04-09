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
	GetOrderForAdmin(ctx context.Context, input model.GetOrderInput) (orderDetail model.GetOrderOutput, result int, err error)
	CreatePayment(ctx context.Context, input model.CreatePaymentInput) (result int, err error)
	GetOrderStatus(ctx context.Context, input model.GetOrderStatusInput) (orderStatus model.GetOrderStatusOutput, result int, err error)
	GetOrderSummary(ctx context.Context) (orderSummary []model.GetOrderSummaryOutput, result int, err error)
}

type orderService struct {
	db          *sql.DB
	cartRepo    repo.ICartRepo
	productRepo repo.IProductRepo
	orderRepo   repo.IOrderRepo
}

// GetOrderSummary implements IOrderService.
func (os *orderService) GetOrderSummary(ctx context.Context) (orderSummary []model.GetOrderSummaryOutput, result int, err error) {
	orderSummaryRepo, err := os.orderRepo.GetOrderSummary(ctx)
	if err != nil {
		return orderSummary, response.ErrCodeInternal, err
	}

	for _, row := range orderSummaryRepo {
		orderSummary = append(orderSummary, model.GetOrderSummaryOutput{
			Status:      string(row.Status.OrdersStatus),
			TotalPrice:  row.TotalPrice,
			TotalAmount: row.TotalAmount,
		})
	}

	return orderSummary, response.ErrCodeSuccess, nil
}

// GetOrderForAdmin implements IOrderService.
func (os *orderService) GetOrderForAdmin(ctx context.Context, input model.GetOrderInput) (orderDetail model.GetOrderOutput, result int, err error) {
	order, err := os.orderRepo.GetOrderForAdmin(ctx, input)
	if err == sql.ErrNoRows {
		return orderDetail, response.ErrCodeOrderNotFound, err
	}

	if err != nil {
		return orderDetail, response.ErrCodeInternal, err
	}

	items, err := os.orderRepo.GetOrderItems(ctx, order.ID)
	if err != nil {
		return orderDetail, response.ErrCodeInternal, err
	}

	orderDetail = model.GetOrderOutput{
		OrderID:          order.ID,
		CreatedAt:        order.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		Status:           string(order.Status.OrdersStatus),
		ShippingAddreess: order.ShippingAddress,
		Payment_method:   string(order.PaymentMethod),
		Total:            order.Total,
	}
	for _, item := range items {
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

// GetOrderStatus implements IOrderService.
func (os *orderService) GetOrderStatus(ctx context.Context, input model.GetOrderStatusInput) (orderStatus model.GetOrderStatusOutput, result int, err error) {
	status, err := os.orderRepo.GetOrderStatus(ctx, input)
	if err == sql.ErrNoRows {
		return orderStatus, response.ErrCodeOrderNotFound, err
	}

	if err != nil {
		return orderStatus, response.ErrCodeInternal, err
	}

	orderStatus.Status = string(status.OrdersStatus)

	return orderStatus, response.ErrCodeSuccess, nil

}

// CreatePayment implements IOrderService.
func (os *orderService) CreatePayment(ctx context.Context, input model.CreatePaymentInput) (result int, err error) {
	//Check order status first
	order, err := os.orderRepo.GetOrder(ctx, model.GetOrderInput{
		OrderID: input.OrderID,
	})
	if err == sql.ErrNoRows {
		return response.ErrCodeOrderNotFound, err
	}

	if err != nil {
		return response.ErrCodeInternal, err
	}

	if order.Status.OrdersStatus != database.OrdersStatusPending {
		return response.ErrCodeStatusNotValid, err
	}

	//Simulate third party payment
	paymentStatus, err := os.orderRepo.CreatePayment(ctx, input)
	if err != nil {
		return response.ErrCodePaymentNotSuccess, err
	}

	//Check result and decide to update order status
	if !paymentStatus {
		return response.ErrCodePaymentNotSuccess, err
	}

	//Update order status
	_, err = os.orderRepo.UpdateStatus(ctx, model.UpdateStatusInput{
		OrderID: input.OrderID,
		Status:  string(database.OrdersStatusPaid),
	})
	if err != nil {
		return response.ErrCodeInternal, err
	}

	return response.ErrCodeSuccess, nil

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
			OrderID:          order.OrderID,
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
			OrderID:          order.ID,
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
	products := make(map[string]database.GetProductByIDForUpdateRow)

	for _, item := range cart {
		product, err := os.productRepo.WithTx(tx).GetProductByIDForUpdate(ctx, item.ProductID)
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
	orderID, err := os.orderRepo.WithTx(tx).CreateOrder(ctx, model.CreateOrderInput{
		PaymentMethod:   input.PaymentMethod,
		ShippingAddress: input.ShippingAddress,
		TotalPrice:      totalPrice,
	})
	if err != nil {
		tx.Rollback()
		return response.ErrCodeInternal, err
	}

	//5. Create order_items
	for _, item := range cart {
		_, err = os.orderRepo.WithTx(tx).CreateOrderItem(ctx, model.CreateOrderItemInput{
			OrderID:     orderID,
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
			ID:       item.ProductID,
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

	items, err := os.orderRepo.GetOrderItems(ctx, order.ID)
	if err != nil {
		return orderDetail, response.ErrCodeInternal, err
	}

	orderDetail = model.GetOrderOutput{
		OrderID:          order.ID,
		CreatedAt:        order.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		Status:           string(order.Status.OrdersStatus),
		ShippingAddreess: order.ShippingAddress,
		Payment_method:   string(order.PaymentMethod),
		Total:            order.Total,
	}
	for _, item := range items {
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
