-- name: CreateProduct :execresult
INSERT INTO `product` (
    name,
    description,
    price,
    quantity,
    image_url
)
VALUES (?, ?, ?, ?, ?)
ON DUPLICATE KEY 
UPDATE
    quantity = quantity  + VALUES(quantity);

-- name: GetProducts :many
SELECT id, name, description, price, quantity, image_url
FROM `product`
WHERE deleted_at IS NULL
LIMIT ?
OFFSET ?;

-- name: GetProductByID :one
SELECT
    id, name, description, price, quantity, image_url
FROM 
    `product`
WHERE
    deleted_at IS NULL AND id = ?;

-- name: UpdateProduct :execresult
UPDATE `product`
SET
    name = ?,
    description = ?,
    price = ?,
    quantity = ?,
    image_url = ?
WHERE id = ?;

-- name: DeleteProduct :execresult
UPDATE `product`
SET deleted_at = NOW()
WHERE id = ?;

-- name: GetQuantity :one
SELECT
    quantity
FROM 
    `product`
WHERE 
    id = ?
