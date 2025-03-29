-- name: CreateOrder :execresult
INSERT INTO `orders`(
    user_id,
    payment_method,
    shipping_address,
    total
)
VALUES(?, ?, ?, ?);