-- name: UpdateUserProfile :execresult
UPDATE `user_profile`
SET
    first_name = ?,
    last_name = ?,
    phone_number = ?,
    address = ?
WHERE user_id = ?;