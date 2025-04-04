package repo

import (
	"context"
	"database/sql"

	"github.com/anle/codebase/internal/database"
	"github.com/anle/codebase/internal/model"
)

type IProductRepo interface {
	CreateProduct(ctx context.Context, input model.CreateProductInput) (result sql.Result, err error)
	UpdateProduct(ctx context.Context, input model.UpdateProductInput) (result sql.Result, err error)
	UpdateProductStatus(ctx context.Context, id int) (result sql.Result, err error)
	UpdateProductByID(ctx context.Context, input model.UpdateProductByIDInput) (result sql.Result, err error)
	DeleteProduct(ctx context.Context, input model.DeleteProductInput) (result sql.Result, err error)
	GetProducts(ctx context.Context, input model.GetProductInput) (products []database.GetProductsRow, err error)
	GetProductByID(ctx context.Context, productID int) (product database.GetProductByIDRow, err error)
	GetProductByIDForUpdate(ctx context.Context, productID int) (product database.GetProductByIDForUpdateRow, err error)
	GetProductForCreate(ctx context.Context, input model.CreateProductInput) (product database.GetProductForCreateRow, err error)
	GetQuantity(ctx context.Context, productID int) (quantity int32, err error)
	WithTx(tx *sql.Tx) IProductRepo
}

type productRepo struct {
	queries *database.Queries
}

// UpdateProductStatus implements IProductRepo.
func (pr *productRepo) UpdateProductStatus(ctx context.Context, id int) (result sql.Result, err error) {
	result, err = pr.queries.UpdateProductStatus(ctx, int32(id))
	if err != nil {
		return result, err
	}

	return result, nil
}

// GetProductForCreate implements IProductRepo.
func (pr *productRepo) GetProductForCreate(ctx context.Context, input model.CreateProductInput) (product database.GetProductForCreateRow, err error) {
	product, err = pr.queries.GetProductForCreate(ctx, database.GetProductForCreateParams{
		Name:  input.Name,
		Price: int64(input.Price),
	})
	if err != nil {
		return product, err
	}

	return product, nil
}

// UpdateProductByID implements IProductRepo.
func (pr *productRepo) UpdateProductByID(ctx context.Context, input model.UpdateProductByIDInput) (result sql.Result, err error) {
	result, err = pr.queries.UpdateQuantity(ctx, database.UpdateQuantityParams{
		ID:       int32(input.ID),
		Quantity: int32(input.Quantity),
	})
	if err != nil {
		return result, err
	}

	return result, nil
}

// WithTx implements IProductRepo.
func (pr *productRepo) WithTx(tx *sql.Tx) IProductRepo {
	return &productRepo{
		queries: pr.queries.WithTx(tx),
	}
}

// GetProductByIDForUpdate implements IProductRepo.
func (pr *productRepo) GetProductByIDForUpdate(ctx context.Context, productID int) (product database.GetProductByIDForUpdateRow, err error) {
	product, err = pr.queries.GetProductByIDForUpdate(ctx, int32(productID))
	if err != nil {
		return product, err
	}

	return product, nil
}

// GetProductByID implements IProductRepo.
func (pr *productRepo) GetProductByID(ctx context.Context, productID int) (product database.GetProductByIDRow, err error) {
	product, err = pr.queries.GetProductByID(ctx, int32(productID))
	if err != nil {
		return product, err
	}

	return product, nil
}

// GetQuantity implements IProductRepo.
func (pr *productRepo) GetQuantity(ctx context.Context, productID int) (quantity int32, err error) {
	quantity, err = pr.queries.GetQuantity(ctx, int32(productID))
	if err != nil {
		return quantity, err
	}

	return quantity, nil
}

// DeleteProduct implements IProductRepo.
func (pr *productRepo) DeleteProduct(ctx context.Context, input model.DeleteProductInput) (result sql.Result, err error) {
	result, err = pr.queries.DeleteProduct(ctx, int32(input.ID))
	if err != nil {
		return result, err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return result, sql.ErrNoRows
	}

	return result, nil
}

// UpdateProduct implements IProductRepo.
func (pr *productRepo) UpdateProduct(ctx context.Context, input model.UpdateProductInput) (result sql.Result, err error) {
	result, err = pr.queries.UpdateProduct(ctx, database.UpdateProductParams{
		ID:          int32(input.ID),
		Name:        input.Name,
		Description: sql.NullString{String: input.Description, Valid: input.Description != ""},
		Price:       int64(input.Price),
		Quantity:    int32(input.Quantity),
		ImageUrl:    input.ImageURL,
	})
	if err != nil {
		return result, err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return result, sql.ErrNoRows
	}

	return result, nil
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
