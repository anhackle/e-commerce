package model

type AddToCartInput struct {
	ProductID string `json:"product_id" binding:"required,uuid"`
	Quantity  int    `json:"quantity" binding:"required,numeric"`
}

type GetCartOutput struct {
	CartID    string `json:"item_id"`
	ProductID string `json:"product_id"`
	Name      string `json:"product_name"`
	Price     int    `json:"price"`
	Quantity  int    `json:"quantity"`
	ImageURL  string `json:"image_url"`
	Total     int    `json:"total"`
}

type DeleteCartInput struct {
	ItemID string `json:"item_id" binding:"required,uuid"`
}

type UpdateCartInput struct {
	ItemID    string `json:"item_id" binding:"required,uuid"`
	ProductID string `json:"product_id" binding:"required,uuid"`
	Quantity  *int   `json:"quantity" binding:"required,numeric,min=0"`
}
