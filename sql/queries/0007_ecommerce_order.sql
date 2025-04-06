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
WHERE id = ? AND user_id = ?;

-- name: GetOrderForAdmin :one
SELECT
    id,
    created_at,
    status,
    shipping_address,
    payment_method,
    total
FROM `orders`
WHERE id = ?;

-- name: UpdateStatus :execresult
UPDATE `orders`
SET status = ?
WHERE id = ?;

-- name: GetOrdersForAdmin :many
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
FROM `orders` o
JOIN `user` u  ON o.user_id = u.id
JOIN `user_profile` up ON u.id = up.user_id
WHERE 
    o.status = IF(? = '', o.status, ?) AND 
    o.payment_method = IF(? = '', o.payment_method, ?)
LIMIT ?
OFFSET ?;

-- name: GetOrderStatus :one
SELECT status
FROM `orders`
WHERE id = ? AND user_id = ?;
