-- name: CreateUser :execresult
INSERT INTO `user` (
    email,
    password
)
VALUES (?, ?);

-- name: FindByEmail :one
SELECT id, email, password, role
FROM `user`
WHERE email = ?;

-- name: CreateUserProfile :execresult
INSERT INTO `user_profile` (
    user_id
) VALUES (?);

-- name: FindByUserId :one
SELECT id, email, password, role
FROM `user`
WHERE id = ?;

-- name: UpdatePassword :execresult
UPDATE `user`
SET
    password = ?
WHERE id = ?;

-- name: UpdateRole :execresult
UPDATE `user`
SET role = ?
WHERE id = ?;