// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: 0005_ecommerce_order.sql

package database

import (
	"context"
	"database/sql"
)

const createOrder = `-- name: CreateOrder :execresult
INSERT INTO ` + "`" + `orders` + "`" + `(
    id,
    user_id,
    payment_method,
    shipping_address,
    total
)
VALUES(?, ?, ?, ?, ?)
`

type CreateOrderParams struct {
	ID              string
	UserID          string
	PaymentMethod   OrdersPaymentMethod
	ShippingAddress string
	Total           int64
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createOrder,
		arg.ID,
		arg.UserID,
		arg.PaymentMethod,
		arg.ShippingAddress,
		arg.Total,
	)
}

const getOrder = `-- name: GetOrder :one
SELECT
    id,
    created_at,
    status,
    shipping_address,
    payment_method,
    total
FROM ` + "`" + `orders` + "`" + `
WHERE id = ? AND user_id = ?
`

type GetOrderParams struct {
	ID     string
	UserID string
}

type GetOrderRow struct {
	ID              string
	CreatedAt       sql.NullTime
	Status          NullOrdersStatus
	ShippingAddress string
	PaymentMethod   OrdersPaymentMethod
	Total           int64
}

func (q *Queries) GetOrder(ctx context.Context, arg GetOrderParams) (GetOrderRow, error) {
	row := q.db.QueryRowContext(ctx, getOrder, arg.ID, arg.UserID)
	var i GetOrderRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Status,
		&i.ShippingAddress,
		&i.PaymentMethod,
		&i.Total,
	)
	return i, err
}

const getOrderForAdmin = `-- name: GetOrderForAdmin :one
SELECT
    id,
    created_at,
    status,
    shipping_address,
    payment_method,
    total
FROM ` + "`" + `orders` + "`" + `
WHERE id = ?
`

type GetOrderForAdminRow struct {
	ID              string
	CreatedAt       sql.NullTime
	Status          NullOrdersStatus
	ShippingAddress string
	PaymentMethod   OrdersPaymentMethod
	Total           int64
}

func (q *Queries) GetOrderForAdmin(ctx context.Context, id string) (GetOrderForAdminRow, error) {
	row := q.db.QueryRowContext(ctx, getOrderForAdmin, id)
	var i GetOrderForAdminRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Status,
		&i.ShippingAddress,
		&i.PaymentMethod,
		&i.Total,
	)
	return i, err
}

const getOrderStatus = `-- name: GetOrderStatus :one
SELECT status
FROM ` + "`" + `orders` + "`" + `
WHERE id = ? AND user_id = ?
`

type GetOrderStatusParams struct {
	ID     string
	UserID string
}

func (q *Queries) GetOrderStatus(ctx context.Context, arg GetOrderStatusParams) (NullOrdersStatus, error) {
	row := q.db.QueryRowContext(ctx, getOrderStatus, arg.ID, arg.UserID)
	var status NullOrdersStatus
	err := row.Scan(&status)
	return status, err
}

const getOrderSummary = `-- name: GetOrderSummary :many
SELECT
    status,
    COUNT(*) AS total_amount,
    CAST(SUM(total) AS SIGNED) AS total_price
FROM ` + "`" + `orders` + "`" + `
GROUP BY status
`

type GetOrderSummaryRow struct {
	Status      NullOrdersStatus
	TotalAmount int64
	TotalPrice  int64
}

func (q *Queries) GetOrderSummary(ctx context.Context) ([]GetOrderSummaryRow, error) {
	rows, err := q.db.QueryContext(ctx, getOrderSummary)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetOrderSummaryRow
	for rows.Next() {
		var i GetOrderSummaryRow
		if err := rows.Scan(&i.Status, &i.TotalAmount, &i.TotalPrice); err != nil {
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

const getOrders = `-- name: GetOrders :many
SELECT 
    id,
    created_at,
    status,
    shipping_address,
    payment_method,
    total
FROM orders
WHERE user_id = ?
ORDER BY created_at DESC
LIMIT ?
OFFSET ?
`

type GetOrdersParams struct {
	UserID string
	Limit  int32
	Offset int32
}

type GetOrdersRow struct {
	ID              string
	CreatedAt       sql.NullTime
	Status          NullOrdersStatus
	ShippingAddress string
	PaymentMethod   OrdersPaymentMethod
	Total           int64
}

func (q *Queries) GetOrders(ctx context.Context, arg GetOrdersParams) ([]GetOrdersRow, error) {
	rows, err := q.db.QueryContext(ctx, getOrders, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetOrdersRow
	for rows.Next() {
		var i GetOrdersRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Status,
			&i.ShippingAddress,
			&i.PaymentMethod,
			&i.Total,
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

const getOrdersForAdmin = `-- name: GetOrdersForAdmin :many
SELECT
    o.user_id AS user_id,
    o.id AS order_id,
    up.first_name AS first_name,
    up.last_name AS last_name,
    up.phone_number AS phone_number,
    o.created_at AS created_at,
    o.status AS status,
    o.shipping_address AS shipping_address,
    o.payment_method AS payment_method,
    o.total AS total
FROM ` + "`" + `orders` + "`" + ` o
JOIN ` + "`" + `user` + "`" + ` u  ON o.user_id = u.id
JOIN ` + "`" + `user_profile` + "`" + ` up ON u.id = up.user_id
WHERE 
    o.status = IF(? = '', o.status, ?) AND 
    o.payment_method = IF(? = '', o.payment_method, ?)
ORDER BY
    CASE WHEN ? = 'created_at' THEN o.created_at
         WHEN ? = 'total' THEN o.total ELSE o.created_at
    END ASC
LIMIT ?
OFFSET ?
`

type GetOrdersForAdminParams struct {
	Column1 interface{}
	IF      interface{}
	Column3 interface{}
	IF_2    interface{}
	Column5 interface{}
	Column6 interface{}
	Limit   int32
	Offset  int32
}

type GetOrdersForAdminRow struct {
	UserID          string
	OrderID         string
	FirstName       sql.NullString
	LastName        sql.NullString
	PhoneNumber     sql.NullString
	CreatedAt       sql.NullTime
	Status          NullOrdersStatus
	ShippingAddress string
	PaymentMethod   OrdersPaymentMethod
	Total           int64
}

func (q *Queries) GetOrdersForAdmin(ctx context.Context, arg GetOrdersForAdminParams) ([]GetOrdersForAdminRow, error) {
	rows, err := q.db.QueryContext(ctx, getOrdersForAdmin,
		arg.Column1,
		arg.IF,
		arg.Column3,
		arg.IF_2,
		arg.Column5,
		arg.Column6,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetOrdersForAdminRow
	for rows.Next() {
		var i GetOrdersForAdminRow
		if err := rows.Scan(
			&i.UserID,
			&i.OrderID,
			&i.FirstName,
			&i.LastName,
			&i.PhoneNumber,
			&i.CreatedAt,
			&i.Status,
			&i.ShippingAddress,
			&i.PaymentMethod,
			&i.Total,
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

const updateStatus = `-- name: UpdateStatus :execresult
UPDATE ` + "`" + `orders` + "`" + `
SET status = ?
WHERE id = ?
`

type UpdateStatusParams struct {
	Status NullOrdersStatus
	ID     string
}

func (q *Queries) UpdateStatus(ctx context.Context, arg UpdateStatusParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateStatus, arg.Status, arg.ID)
}
