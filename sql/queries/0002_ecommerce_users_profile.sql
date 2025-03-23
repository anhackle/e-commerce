-- name: UpdateUserProfile :execresult
UPDATE `user_profile`
SET
    first_name = ?,
    last_name = ?,
    phone_number = ?,
    address = ?
WHERE user_id = ?;

-- name: GetUserProfile :one
SELECT first_name, last_name, phone_number, address
FROM `user_profile`
WHERE user_id = ?;