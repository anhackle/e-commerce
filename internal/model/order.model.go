package model

type CreateOrderInput struct {
	PaymentMethod   string `json:"payment_method" binding:"required,paymentmethod"`
	ShippingAddress string `json:"shipping_address" binding:"required"`
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
