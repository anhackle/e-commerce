package model

type CreateOrderInput struct {
	PaymentMethod   string `json:"payment_method" binding:"required,paymentmethod"`
	ShippingAddress string `json:"shipping_address" binding:"required,endsnotwith= ,startsnotwith= "`
	TotalPrice      int64  `json:"total_price"`
}

// Internal input => Not need binding
type CreateOrderItemInput struct {
	OrderID     int32
	Name        string
	Description string
	Price       int64
	Quantity    int32
	ImageUrl    string
}

type GetOrdersInput struct {
	Limit int `json:"limit" binding:"required,numeric,max=20,gt=0"`
	Page  int `json:"page" binding:"required,numeric,gt=0"`
}

type GetOrderInput struct {
	OrderID int `json:"order_id" binding:"required,numeric,min=0"`
}

type GetOrdersOutput struct {
	OrderID          int    `json:"order_id"`
	CreatedAt        string `json:"created_at"`
	Status           string `json:"status"`
	ShippingAddreess string `json:"shipping_address"`
	Payment_method   string `json:"payment_method"`
	Total            int64  `json:"total"`
}

type GetOrderOutput struct {
	OrderID          int
	CreatedAt        string
	Status           string
	ShippingAddreess string
	Payment_method   string
	Total            int64
	Items            []GetOrderItemsOutput
}

type GetOrderItemsOutput struct {
	Name        string
	Description string
	Price       int64
	Quantity    int
	Image_url   string
}

type UpdateStatusInput struct {
	OrderID int    `json:"order_id" binding:"required,numeric,min=0"`
	Status  string `json:"status" binding:"required,status"`
}

type GetOrdersForAdminInput struct {
	Limit          int    `json:"limit" binding:"required,numeric,max=20,gt=0"`
	Page           int    `json:"page" binding:"required,numeric,gt=0"`
	Status         string `json:"status"`
	Payment_method string `json:"payment_method"`
	OrderBy        string `json:"order_by" binding:"order_by"`
}

type GetOrdersForAdminOutput struct {
	OrderID          int    `json:"order_id"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	PhoneNumber      string `json:"phone_number"`
	CreatedAt        string `json:"created_at"`
	Status           string `json:"status"`
	ShippingAddreess string `json:"shipping_address"`
	Payment_method   string `json:"payment_method"`
	Total            int64  `json:"total"`
}

type CreatePaymentInput struct {
	OrderID int `json:"order_id" binding:"required,numeric,min=0"`
}

type GetOrderStatusInput struct {
	OrderID int `json:"order_id" binding:"required,numeric,min=0"`
}

type GetOrderStatusOutput struct {
	Status string `json:"status"`
}

type GetOrderSummaryOutput struct {
	Status      string `json:"status"`
	TotalPrice  int64  `json:"total_price"`
	TotalAmount int64  `json:"total_amount"`
}
