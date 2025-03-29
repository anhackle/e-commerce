// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: 0007_ecommerce_order.sql

package database

import (
	"context"
	"database/sql"
)

const createOrder = `-- name: CreateOrder :execresult
INSERT INTO ` + "`" + `order` + "`" + `(
    user_id,
    payment_method,
    shipping_address,
    total
)
VALUES(?, ?, ?, ?)
`

type CreateOrderParams struct {
	UserID          int32
	PaymentMethod   OrderPaymentMethod
	ShippingAddress string
	Total           int64
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createOrder,
		arg.UserID,
		arg.PaymentMethod,
		arg.ShippingAddress,
		arg.Total,
	)
}
