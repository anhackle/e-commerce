-- name: CreateProduct :execresult
INSERT INTO `product` (
    name,
    description,
    price,
    quantity,
    image_url
)
VALUES (?, ?, ?, ?, ?);

-- name: GetProducts :many
SELECT id, name, description, price, quantity, image_url
FROM `product`
LIMIT ?
OFFSET ?;
