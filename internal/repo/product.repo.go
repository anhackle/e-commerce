package repo

import (
	"context"
	"database/sql"

	"github.com/anle/codebase/internal/database"
	"github.com/anle/codebase/internal/model"
)

type IProductRepo interface {
	CreateProduct(ctx context.Context, input model.CreateProductInput) (result sql.Result, err error)
	GetProducts(ctx context.Context, input model.GetProductInput) (products []database.GetProductsRow, err error)
}

type productRepo struct {
	queries *database.Queries
}

// GetProducts implements IProductRepo.
func (pr *productRepo) GetProducts(ctx context.Context, input model.GetProductInput) (products []database.GetProductsRow, err error) {
	products, err = pr.queries.GetProducts(ctx, database.GetProductsParams{
		Limit:  int32(input.Limit),
		Offset: int32(input.Page),
	})
	if err != nil {
		return products, err
	}

	return products, nil
}

// CreateProduct implements IProductRepo.
func (pr *productRepo) CreateProduct(ctx context.Context, input model.CreateProductInput) (result sql.Result, err error) {
	result, err = pr.queries.CreateProduct(ctx, database.CreateProductParams{
		Name:        input.Name,
		Description: sql.NullString{String: input.Description, Valid: input.Description != ""},
		Price:       int64(input.Price),
		Quantity:    int32(input.Quantity),
		ImageUrl:    input.ImageURL,
	})
	if err != nil {
		return result, err
	}

	return result, nil
}

func NewProductRepo(dbConn *sql.DB) IProductRepo {
	return &productRepo{
		queries: database.New(dbConn),
	}
}
