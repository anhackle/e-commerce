package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/anle/codebase/internal/model"
	"github.com/anle/codebase/internal/repo"
	"github.com/anle/codebase/response"
)

type ICartService interface {
	AddToCart(ctx context.Context, input model.AddToCartInput) (result int, err error)
	GetCart(ctx context.Context) (cart []model.GetCartOutput, result int, err error)
	DeleteCart(ctx context.Context, input model.DeleteCartInput) (result int, err error)
}

type cartService struct {
	cartRepo    repo.ICartRepo
	productRepo repo.IProductRepo
}

// DeleteCart implements ICartService.
func (cs *cartService) DeleteCart(ctx context.Context, input model.DeleteCartInput) (result int, err error) {
	_, err = cs.cartRepo.DeleteCart(ctx, input)
	if err != nil && err == sql.ErrNoRows {
		return response.ErrCodeItemNotFoundInCart, err
	}

	if err != nil && err != sql.ErrNoRows {
		return response.ErrCodeInternal, err
	}

	return response.ErrCodeSuccess, nil
}

// AddToCart implements ICartService.
func (cs *cartService) AddToCart(ctx context.Context, input model.AddToCartInput) (result int, err error) {
	_, err = cs.productRepo.GetProductByID(ctx, input.ProductID)

	if err != nil && err == sql.ErrNoRows {
		return response.ErrCodeProductNotFound, err
	}

	if err != nil && err != sql.ErrNoRows {
		return response.ErrCodeInternal, err
	}

	stockQuantity, err := cs.productRepo.GetQuantity(ctx, input.ProductID)
	if err != nil {
		return response.ErrCodeInternal, nil
	}

	if input.Quantity > int(stockQuantity) {
		return response.ErrCodeQuantityNotEnough, errors.New("quantity not enough")
	}

	_, err = cs.cartRepo.AddToCart(ctx, input)
	if err != nil {
		return response.ErrCodeInternal, err
	}

	return response.ErrCodeSuccess, nil
}

// GetCart implements ICartService.
func (cs *cartService) GetCart(ctx context.Context) (cart []model.GetCartOutput, result int, err error) {
	cartItems, err := cs.cartRepo.GetCart(ctx)
	if err != nil {
		return []model.GetCartOutput{}, response.ErrCodeInternal, err
	}

	for _, item := range cartItems {
		cart = append(cart, model.GetCartOutput{
			CartID:    int(item.CartID),
			ProductID: int(item.ProductID),
			Name:      item.ProductName,
			Price:     int(item.ProductPrice),
			Quantity:  int(item.Quantity),
			ImageURL:  item.ImageUrl,
			Total:     int(item.Total),
		})
	}

	return cart, response.ErrCodeSuccess, nil
}

func NewCartService(cartRepo repo.ICartRepo, productRepo repo.IProductRepo) ICartService {
	return &cartService{
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}
