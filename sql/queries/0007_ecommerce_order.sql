-- name: CreateOrder :execresult
INSERT INTO `orders`(
    user_id,
    payment_method,
    shipping_address,
    total
)
VALUES(?, ?, ?, ?);

-- name: GetOrders :many
SELECT 
    id,
    created_at,
    status,
    shipping_address,
    payment_method,
    total
FROM orders
WHERE user_id = ?
LIMIT ?
OFFSET ?;

-- name: GetOrder :one
SELECT
    id,
    created_at,
    status,
    shipping_address,
    payment_method,
    total
FROM `orders`
WHERE id = ? AND user_id = ?