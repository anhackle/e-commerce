-- name: CreateUser :execresult
INSERT INTO `user` (
    email,
    password
)
VALUES (?, ?);

-- name: FindByEmail :one
SELECT id, email, password
FROM `user`
WHERE email = ?;

-- name: CreateUserProfile :execresult
INSERT INTO `user_profile` (
    user_id
) VALUES (?);

-- name: FindByUserId :one
SELECT id, email, password
FROM `user`
WHERE id = ?;

-- name: UpdatePassword :execresult
UPDATE `user`
SET
    password = ?
WHERE id = ?