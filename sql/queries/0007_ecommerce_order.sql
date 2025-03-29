-- name: CreateOrder :execresult
INSERT INTO `order`(
    user_id,
    payment_method,
    shipping_address,
    total
)
VALUES(?, ?, ?, ?);