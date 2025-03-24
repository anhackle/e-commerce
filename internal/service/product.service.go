package service

import (
	"context"

	"github.com/anle/codebase/internal/model"
	"github.com/anle/codebase/internal/repo"
	"github.com/anle/codebase/response"
)

type IProductService interface {
	CreateProduct(ctx context.Context, input model.CreateProductInput) (result int, err error)
	UpdateProduct(ctx context.Context, input model.UpdateProductInput) (result int, err error)
	DeleteProduct(ctx context.Context, input model.DeleteProductInput) (result int, err error)
	GetProducts(ctx context.Context, input model.GetProductInput) (products []model.GetProductOutput, result int, err error)
}

type productService struct {
	productRepo repo.IProductRepo
}

// DeleteProduct implements IProductService.
func (ps *productService) DeleteProduct(ctx context.Context, input model.DeleteProductInput) (result int, err error) {
	_, err = ps.productRepo.DeleteProduct(ctx, input)
	if err != nil {
		return response.ErrCodeInternal, err
	}

	return response.ErrCodeSuccess, nil
}

// UpdateProduct implements IProductService.
func (ps *productService) UpdateProduct(ctx context.Context, input model.UpdateProductInput) (result int, err error) {
	_, err = ps.productRepo.UpdateProduct(ctx, input)
	if err != nil {
		return response.ErrCodeInternal, err
	}

	return response.ErrCodeSuccess, nil
}

// GetProducts implements IProductService.
func (ps *productService) GetProducts(ctx context.Context, input model.GetProductInput) (products []model.GetProductOutput, result int, err error) {
	var getProductInput = model.GetProductInput{
		Limit: input.Limit,
		Page:  (input.Page - 1) * input.Limit,
	}
	productsRepo, err := ps.productRepo.GetProducts(ctx, getProductInput)
	if err != nil {
		return products, response.ErrCodeInternal, err
	}

	for _, product := range productsRepo {
		products = append(products, model.GetProductOutput{
			ID:          int(product.ID),
			Name:        product.Name,
			Description: product.Description.String,
			Price:       int(product.Price),
			Quantity:    int(product.Quantity),
			ImageURL:    product.ImageUrl,
		})
	}

	return products, response.ErrCodeSuccess, nil
}

// CreateProduct implements IProductService.
func (ps *productService) CreateProduct(ctx context.Context, input model.CreateProductInput) (result int, err error) {
	_, err = ps.productRepo.CreateProduct(ctx, input)
	if err != nil {
		return response.ErrCodeInternal, err
	}

	return response.ErrCodeSuccess, nil
}

func NewProductService(productRepo repo.IProductRepo) IProductService {
	return &productService{
		productRepo: productRepo,
	}
}
