-- name: CreateOrderItem :execresult
INSERT INTO `order_item`(
    id,
    order_id,
    name,
    description,
    price,
    quantity,
    image_url
)
VALUES(?, ?, ?, ?, ?, ?, ?);

-- name: GetOrderItems :many
SELECT
    name,
    description,
    price,
    quantity,
    image_url
FROM `order_item`
WHERE order_id = ?