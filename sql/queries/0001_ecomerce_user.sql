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