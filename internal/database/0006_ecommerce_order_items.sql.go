// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: 0006_ecommerce_order_items.sql

package database

import (
	"context"
	"database/sql"
)

const createOrderItem = `-- name: CreateOrderItem :execresult
INSERT INTO ` + "`" + `order_item` + "`" + `(
    id,
    order_id,
    name,
    description,
    price,
    quantity,
    image_url
)
VALUES(?, ?, ?, ?, ?, ?, ?)
`

type CreateOrderItemParams struct {
	ID          string
	OrderID     string
	Name        string
	Description sql.NullString
	Price       int64
	Quantity    int32
	ImageUrl    string
}

func (q *Queries) CreateOrderItem(ctx context.Context, arg CreateOrderItemParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createOrderItem,
		arg.ID,
		arg.OrderID,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.Quantity,
		arg.ImageUrl,
	)
}

const getOrderItems = `-- name: GetOrderItems :many
SELECT
    name,
    description,
    price,
    quantity,
    image_url
FROM ` + "`" + `order_item` + "`" + `
WHERE order_id = ?
`

type GetOrderItemsRow struct {
	Name        string
	Description sql.NullString
	Price       int64
	Quantity    int32
	ImageUrl    string
}

func (q *Queries) GetOrderItems(ctx context.Context, orderID string) ([]GetOrderItemsRow, error) {
	rows, err := q.db.QueryContext(ctx, getOrderItems, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetOrderItemsRow
	for rows.Next() {
		var i GetOrderItemsRow
		if err := rows.Scan(
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
