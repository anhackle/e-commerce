package repo

import (
	"context"
	"database/sql"

	"github.com/anle/codebase/internal/dao"
	"github.com/anle/codebase/internal/database"
	uuidv4 "github.com/anle/codebase/internal/utils/uuid"

	"github.com/anle/codebase/internal/model"
)

type IProductRepo interface {
	CreateProduct(ctx context.Context, input model.CreateProductInput) (productID string, err error)
	UpdateProduct(ctx context.Context, input model.UpdateProductInput) (result sql.Result, err error)
	UpdateProductStatus(ctx context.Context, productID string) (result sql.Result, err error)
	UpdateProductByID(ctx context.Context, input model.UpdateProductByIDInput) (result sql.Result, err error)
	DeleteProduct(ctx context.Context, input model.DeleteProductInput) (result sql.Result, err error)
	GetProducts(ctx context.Context, input model.GetProductsInput) (products []database.GetProductsRow, err error)
	GetProductsWithSearchForAdmin(ctx context.Context, input model.GetProductsForAdminInput) (products []dao.GetProductsWithSearchForAdminRow, err error)
	GetProductByID(ctx context.Context, productID string) (product database.GetProductByIDRow, err error)
	GetProductByIDForUpdate(ctx context.Context, productID string) (product database.GetProductByIDForUpdateRow, err error)
	GetProductForCreate(ctx context.Context, input model.CreateProductInput) (product database.GetProductForCreateRow, err error)
	GetQuantity(ctx context.Context, productID string) (quantity int32, err error)
	WithTx(tx *sql.Tx) IProductRepo
}

type productRepo struct {
	queries *database.Queries
	dto     *dao.Queries
}

// GetProductsForAdmin implements IProductRepo.
func (pr *productRepo) GetProductsWithSearchForAdmin(ctx context.Context, input model.GetProductsForAdminInput) (products []dao.GetProductsWithSearchForAdminRow, err error) {
	products, err = pr.dto.GetProductsWithSearchForAdmin(ctx, dao.GetProductsWithSearchForAdminParams{
		Limit:     int32(input.Limit),
		Offset:    int32(input.Page),
		FromPrice: input.MinPrice,
		ToPrice:   input.MaxPrice,
		Search:    input.Search,
	})
	if err != nil {
		return products, err
	}

	return products, nil
}

// UpdateProductStatus implements IProductRepo.
func (pr *productRepo) UpdateProductStatus(ctx context.Context, productID string) (result sql.Result, err error) {
	result, err = pr.queries.UpdateProductStatus(ctx, productID)
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
		ID:       input.ID,
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
func (pr *productRepo) GetProductByIDForUpdate(ctx context.Context, productID string) (product database.GetProductByIDForUpdateRow, err error) {
	product, err = pr.queries.GetProductByIDForUpdate(ctx, productID)
	if err != nil {
		return product, err
	}

	return product, nil
}

// GetProductByID implements IProductRepo.
func (pr *productRepo) GetProductByID(ctx context.Context, productID string) (product database.GetProductByIDRow, err error) {
	product, err = pr.queries.GetProductByID(ctx, productID)
	if err != nil {
		return product, err
	}

	return product, nil
}

// GetQuantity implements IProductRepo.
func (pr *productRepo) GetQuantity(ctx context.Context, productID string) (quantity int32, err error) {
	quantity, err = pr.queries.GetQuantity(ctx, (productID))
	if err != nil {
		return quantity, err
	}

	return quantity, nil
}

// DeleteProduct implements IProductRepo.
func (pr *productRepo) DeleteProduct(ctx context.Context, input model.DeleteProductInput) (result sql.Result, err error) {
	result, err = pr.queries.DeleteProduct(ctx, input.ID)
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
		ID:          input.ID,
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
func (pr *productRepo) GetProducts(ctx context.Context, input model.GetProductsInput) (products []database.GetProductsRow, err error) {
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
func (pr *productRepo) CreateProduct(ctx context.Context, input model.CreateProductInput) (productID string, err error) {
	productID = uuidv4.GenerateUUID()
	_, err = pr.queries.CreateProduct(ctx, database.CreateProductParams{
		ID:          productID,
		Name:        input.Name,
		Description: sql.NullString{String: input.Description, Valid: input.Description != ""},
		Price:       int64(input.Price),
		Quantity:    int32(input.Quantity),
		ImageUrl:    input.ImageURL,
	})
	if err != nil {
		return productID, err
	}

	return productID, nil
}

func NewProductRepo(dbConn *sql.DB) IProductRepo {
	return &productRepo{
		queries: database.New(dbConn),
		dto:     dao.New(dbConn),
	}
}
