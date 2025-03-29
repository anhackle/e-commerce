-- name: CreateOrderItem :execresult
INSERT INTO `order_item`(
    order_id,
    name,
    description,
    price,
    quantity,
    image_url
)
VALUES(?, ?, ?, ?, ?, ?);