package service

import (
	"context"
	"database/sql"

	"github.com/anle/codebase/internal/model"
	"github.com/anle/codebase/internal/repo"
	"github.com/anle/codebase/response"
)

type IProductService interface {
	CreateProduct(ctx context.Context, input model.CreateProductInput) (result int, err error)
	UpdateProduct(ctx context.Context, input model.UpdateProductInput) (result int, err error)
	DeleteProduct(ctx context.Context, input model.DeleteProductInput) (result int, err error)
	GetProducts(ctx context.Context, input model.GetProductsInput) (products []model.GetProductsOutput, result int, err error)
	GetProductsForAdmin(ctx context.Context, input model.GetProductsForAdminInput) (products []model.GetProductsForAdminOutput, result int, err error)
}

type productService struct {
	productRepo repo.IProductRepo
}

// GetProductsForAdmin implements IProductService.
func (ps *productService) GetProductsForAdmin(ctx context.Context, input model.GetProductsForAdminInput) (products []model.GetProductsForAdminOutput, result int, err error) {
	input.Page = (input.Page - 1) * input.Limit
	productsRepo, err := ps.productRepo.GetProductsForAdmin(ctx, input)
	if err != nil {
		return products, response.ErrCodeInternal, err
	}

	for _, product := range productsRepo {
		products = append(products, model.GetProductsForAdminOutput{
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

// DeleteProduct implements IProductService.
func (ps *productService) DeleteProduct(ctx context.Context, input model.DeleteProductInput) (result int, err error) {
	_, err = ps.productRepo.DeleteProduct(ctx, input)
	if err == sql.ErrNoRows {
		return response.ErrCodeProductNotFound, err
	}

	if err != nil && err != sql.ErrNoRows {
		return response.ErrCodeInternal, err
	}

	return response.ErrCodeSuccess, nil
}

// UpdateProduct implements IProductService.
func (ps *productService) UpdateProduct(ctx context.Context, input model.UpdateProductInput) (result int, err error) {
	_, err = ps.productRepo.UpdateProduct(ctx, input)
	if err == sql.ErrNoRows {
		return response.ErrCodeProductNotFound, err
	}

	if err != nil && err != sql.ErrNoRows {
		return response.ErrCodeInternal, err
	}

	return response.ErrCodeSuccess, nil
}

// GetProducts implements IProductService.
func (ps *productService) GetProducts(ctx context.Context, input model.GetProductsInput) (products []model.GetProductsOutput, result int, err error) {
	var getProductInput = model.GetProductsInput{
		Limit: input.Limit,
		Page:  (input.Page - 1) * input.Limit,
	}
	productsRepo, err := ps.productRepo.GetProducts(ctx, getProductInput)
	if err != nil {
		return products, response.ErrCodeInternal, err
	}

	for _, product := range productsRepo {
		products = append(products, model.GetProductsOutput{
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
	// Check if product(name,price) exists
	product, err := ps.productRepo.GetProductForCreate(ctx, input)
	if err == sql.ErrNoRows {
		_, err = ps.productRepo.CreateProduct(ctx, input)
		if err != nil {
			return response.ErrCodeInternal, err
		}

		return response.ErrCodeSuccess, nil
	}

	if err != nil {
		return response.ErrCodeInternal, err
	}

	// if product was not deleted --> update quantity = quantity + plus
	if product.DeletedAt.Valid == false {
		result, err = ps.UpdateProduct(ctx, model.UpdateProductInput{
			ID:          int(product.ID),
			Name:        input.Name,
			Description: input.Description,
			Price:       input.Price,
			Quantity:    int(product.Quantity) + input.Quantity,
			ImageURL:    input.ImageURL,
		})

		return result, err
	}

	// if product was deleted --> update deleted_at nul and quantity
	// revive product
	_, err = ps.productRepo.UpdateProductStatus(ctx, int(product.ID))
	if err != nil {
		return response.ErrCodeInternal, err
	}

	result, err = ps.UpdateProduct(ctx, model.UpdateProductInput{
		ID:          int(product.ID),
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Quantity:    input.Quantity,
		ImageURL:    input.ImageURL,
	})

	return result, err
}

func NewProductService(productRepo repo.IProductRepo) IProductService {
	return &productService{
		productRepo: productRepo,
	}
}
