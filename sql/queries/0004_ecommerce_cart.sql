-- name: AddToCart :execresult
INSERT INTO `cart` (
    id,
    user_id,
    product_id,
    quantity
)
VALUES (?, ?, ?, ?)
ON DUPLICATE KEY 
UPDATE
    quantity = quantity  + VALUES(quantity);

-- name: GetCart :many
SELECT
    c.id AS cart_id,
    p.id AS product_id,
    p.name AS product_name,
    c.quantity AS quantity,
    p.image_url AS image_url,
    p.price AS product_price
FROM `cart` c
JOIN
    `product` p ON c.product_id = p.id
WHERE
    c.user_id = ?;

-- name: DeleteCartByID :execresult
DELETE
FROM `cart`
WHERE
    user_id = ? AND id = ?;

-- name: DeleteCart :execresult
DELETE 
FROM `cart`
WHERE
    user_id = ?;

-- name: UpdateCart :execresult
UPDATE `cart`
SET quantity = ?
WHERE
    id = ? AND user_id = ? AND product_id = ?
