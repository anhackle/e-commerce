package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

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
	GetProduct(ctx context.Context, input model.GetProductInput) (product model.GetProductOutput, result int, err error)
}

type productService struct {
	productRepo repo.IProductRepo
	redisCache  repo.IRedisCache
	localCache  repo.ILocalCache
}

// GetProductsForAdmin implements IProductService.
func (ps *productService) GetProductsForAdmin(ctx context.Context, input model.GetProductsForAdminInput) (products []model.GetProductsForAdminOutput, result int, err error) {
	input.Page = (input.Page - 1) * input.Limit
	productsRepo, err := ps.productRepo.GetProductsWithSearchForAdmin(ctx, input)
	if err != nil {
		return products, response.ErrCodeInternal, err
	}

	for _, product := range productsRepo {
		products = append(products, model.GetProductsForAdminOutput{
			ID:          product.ID,
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
			ID:          product.ID,
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
	if !product.DeletedAt.Valid {
		result, err = ps.UpdateProduct(ctx, model.UpdateProductInput{
			ID:          product.ID,
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
	_, err = ps.productRepo.UpdateProductStatus(ctx, product.ID)
	if err != nil {
		return response.ErrCodeInternal, err
	}

	result, err = ps.UpdateProduct(ctx, model.UpdateProductInput{
		ID:          product.ID,
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Quantity:    input.Quantity,
		ImageURL:    input.ImageURL,
	})

	return result, err
}

func (ps *productService) GetKeyProductCache(productID string) string {
	return fmt.Sprintf("product:%s", productID)
}

func (ps *productService) GetProduct(ctx context.Context, input model.GetProductInput) (product model.GetProductOutput, result int, err error) {
	productCache, found := ps.localCache.Get(ctx, ps.GetKeyProductCache(input.ID))
	if found {
		err = json.Unmarshal([]byte(productCache), &product)
		if err != nil {
			return product, response.ErrCodeInternal, err
		}

		return product, response.ErrCodeSuccess, nil
	}

	productCache, err = ps.redisCache.Get(ctx, ps.GetKeyProductCache(input.ID))
	if err != nil {
		return product, response.ErrCodeInternal, err
	}

	if productCache != "" {
		err = json.Unmarshal([]byte(productCache), &product)
		if err != nil {
			return product, response.ErrCodeInternal, err
		}
		ps.localCache.SetWithTTL(ctx, ps.GetKeyProductCache(input.ID), product)

		return product, response.ErrCodeSuccess, nil
	}

	lockey := fmt.Sprintf("lock:productItem:%s", input.ID)
	err = ps.redisCache.WithDistributedLock(ctx, lockey, 5, func(ctx context.Context) error {
		productRepo, err := ps.productRepo.GetProductByID(ctx, input.ID)
		if err != nil {
			return err
		}

		product = model.GetProductOutput{
			ID:          productRepo.ID,
			Name:        productRepo.Name,
			Description: productRepo.Description.String,
			Price:       int(productRepo.Price),
			Quantity:    int(productRepo.Quantity),
			ImageURL:    productRepo.ImageUrl,
		}
		err = ps.redisCache.Set(ctx, ps.GetKeyProductCache(input.ID), product, 60*60)
		if err != nil {
			return err
		}

		_, err = ps.localCache.SetWithTTL(ctx, ps.GetKeyProductCache(input.ID), product)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return product, response.ErrCodeInternal, err
	}

	return product, response.ErrCodeSuccess, nil
}

func NewProductService(productRepo repo.IProductRepo, redisCache repo.IRedisCache, localCache repo.ILocalCache) IProductService {
	return &productService{
		productRepo: productRepo,
		redisCache:  redisCache,
		localCache:  localCache,
	}
}
