package model

type AddToCartInput struct {
	ProductID int `json:"product_id" binding:"required,numeric,min=0"`
	Quantity  int `json:"quantity" binding:"required,numeric"`
}

type GetCartOutput struct {
	CartID    int    `json:"cart_id"`
	ProductID int    `json:"product_id"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Quantity  int    `json:"quantity"`
	ImageURL  string `json:"image_url"`
	Total     int    `json:"total"`
}

type DeleteCartInput struct {
	ItemID int `json:"item_id" binding:"required,numeric,min=0"`
}

type UpdateCartInput struct {
	ItemID    int  `json:"item_id" binding:"required,numeric,min=0"`
	ProductID int  `json:"product_id" binding:"required,numeric,min=0"`
	Quantity  *int `json:"quantity" binding:"required,numeric,min=0"`
}
