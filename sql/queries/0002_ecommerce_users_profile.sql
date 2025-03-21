-- name: CreateUserProfile :execresult
INSERT INTO `user_profile` (
    user_id
) VALUES (?);

-- name: UpdateUserProfile :execresult
UPDATE `user_profile`
SET
    first_name = ?,
    last_name = ?,
    phone_number = ?,
    address = ?
WHERE user_id = ?;