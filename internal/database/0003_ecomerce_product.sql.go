// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: 0003_ecomerce_product.sql

package database

import (
	"context"
	"database/sql"
)

const createProduct = `-- name: CreateProduct :execresult
INSERT INTO ` + "`" + `product` + "`" + ` (
    name,
    description,
    price,
    quantity,
    image_url
)
VALUES (?, ?, ?, ?, ?)
`

type CreateProductParams struct {
	Name        string
	Description sql.NullString
	Price       int64
	Quantity    int32
	ImageUrl    string
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createProduct,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.Quantity,
		arg.ImageUrl,
	)
}

const deleteProduct = `-- name: DeleteProduct :execresult
UPDATE ` + "`" + `product` + "`" + `
SET deleted_at = NOW()
WHERE id = ?
`

func (q *Queries) DeleteProduct(ctx context.Context, id int32) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteProduct, id)
}

const getProducts = `-- name: GetProducts :many
SELECT id, name, description, price, quantity, image_url
FROM ` + "`" + `product` + "`" + `
WHERE deleted_at IS NULL
LIMIT ?
OFFSET ?
`

type GetProductsParams struct {
	Limit  int32
	Offset int32
}

type GetProductsRow struct {
	ID          int32
	Name        string
	Description sql.NullString
	Price       int64
	Quantity    int32
	ImageUrl    string
}

func (q *Queries) GetProducts(ctx context.Context, arg GetProductsParams) ([]GetProductsRow, error) {
	rows, err := q.db.QueryContext(ctx, getProducts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetProductsRow
	for rows.Next() {
		var i GetProductsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.Quantity,
			&i.ImageUrl,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProduct = `-- name: UpdateProduct :execresult
UPDATE ` + "`" + `product` + "`" + `
SET
    name = ?,
    description = ?,
    price = ?,
    quantity = ?,
    image_url = ?
WHERE id = ?
`

type UpdateProductParams struct {
	Name        string
	Description sql.NullString
	Price       int64
	Quantity    int32
	ImageUrl    string
	ID          int32
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateProduct,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.Quantity,
		arg.ImageUrl,
		arg.ID,
	)
}
