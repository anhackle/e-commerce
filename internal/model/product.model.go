package model

type CreateProductInput struct {
	Name        string `json:"name" binding:"required,name,max=255,endsnotwith= ,startsnotwith= "`
	Description string `json:"description" binding:"required,endsnotwith= ,startsnotwith= "`
	Price       int    `json:"price" binding:"required,numeric,min=0"`
	Quantity    int    `json:"quantity" binding:"required,numeric,min=0,max=100"`
	ImageURL    string `json:"image_url" binding:"required,url"`
}

type UpdateProductInput struct {
	ID          int    `json:"product_id" binding:"required,numeric,min=0"`
	Name        string `json:"name" binding:"required,name,max=255,endsnotwith= ,startsnotwith= "`
	Description string `json:"description" binding:"required,endsnotwith= ,startsnotwith= "`
	Price       int    `json:"price" binding:"required,numeric,min=0"`
	Quantity    int    `json:"quantity" binding:"required,numeric,min=0,max=100"`
	ImageURL    string `json:"image_url" binding:"required,url"`
}

// Internal input => Not need binding
type UpdateProductByIDInput struct {
	ID       int
	Quantity int
}

type DeleteProductInput struct {
	ID int `json:"product_id" binding:"required,numeric,min=0"`
}

type GetProductInput struct {
	Limit int `json:"limit" binding:"required,numeric,max=20,gt=0"`
	Page  int `json:"page" binding:"required,numeric,gt=0"`
}

type GetProductOutput struct {
	ID          int    `json:"product_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
	ImageURL    string `json:"image_url"`
}
