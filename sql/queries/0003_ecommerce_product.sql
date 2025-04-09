-- name: CreateProduct :execresult
INSERT INTO `product` (
    id,
    name,
    description,
    price,
    quantity,
    image_url
)
VALUES (?, ?, ?, ?, ?, ?);

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
    id = ? AND deleted_at IS NULL;

-- name: GetProductByIDForUpdate :one
SELECT
    id, name, description, price, quantity, image_url
FROM 
    `product`
WHERE
    id = ? AND deleted_at IS NULL
FOR UPDATE;

-- name: GetQuantity :one
SELECT
    quantity
FROM 
    `product`
WHERE 
    id = ? AND deleted_at IS NULL;

-- name: GetProductForCreate :one
SELECT id, quantity, deleted_at
FROM `product`
WHERE name = ? AND price = ?;

-- name: UpdateProductStatus :execresult
UPDATE `product`
SET
    deleted_at = NULL
WHERE id = ? AND deleted_at IS NOT NULL;

-- name: UpdateProduct :execresult
UPDATE `product`
SET
    name = ?,
    description = ?,
    price = ?,
    quantity = ?,
    image_url = ?
WHERE id = ? AND deleted_at IS NULL;

-- name: UpdateQuantity :execresult
UPDATE `product`
SET
    quantity = ?
WHERE
    id = ? AND deleted_at IS NULL;

-- name: DeleteProduct :execresult
UPDATE `product`
SET deleted_at = NOW()
WHERE id = ? AND deleted_at is NULL;