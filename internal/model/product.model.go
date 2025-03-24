package model

type CreateProductInput struct {
	Name        string `json:"name" binding:"required,name,max=255"`
	Description string `json:"description" binding:"required"`
	Price       int    `json:"price" binding:"required,numeric"`
	Quantity    int    `json:"quantity" binding:"required,numeric"`
	ImageURL    string `json:"image_url" binding:"required,url"`
}

type UpdateProductInput struct {
	ID          int    `json:"id" binding:"required,numeric"`
	Name        string `json:"name" binding:"required,name,max=255"`
	Description string `json:"description" binding:"required"`
	Price       int    `json:"price" binding:"required,numeric"`
	Quantity    int    `json:"quantity" binding:"required,numeric"`
	ImageURL    string `json:"image_url" binding:"required,url"`
}

type DeleteProductInput struct {
	ID int `json:"id" binding:"required,numeric"`
}

type GetProductInput struct {
	Limit int `json:"limit" binding:"required,numeric,max=20,gt=0"`
	Page  int `json:"page" binding:"required,numeric,gt=0"`
}

type GetProductOutput struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
	ImageURL    string `json:"imag_url"`
}
