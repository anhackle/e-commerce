-- name: CreateUser :execresult
INSERT INTO `user` (
    id,
    email,
    password
)
VALUES (?, ?, ?);

-- name: FindByEmail :one
SELECT id, email, password, role
FROM `user`
WHERE email = ? AND deleted_at IS NULL;

-- name: CreateUserProfile :execresult
INSERT INTO `user_profile` (
    id, user_id
) VALUES (?, ?);

-- name: FindByUserId :one
SELECT id, email, password, role
FROM `user`
WHERE id = ? AND deleted_at IS NULL;

-- name: UpdatePassword :execresult
UPDATE `user`
SET
    password = ?
WHERE id = ? AND deleted_at IS NULL;

-- name: UpdateRole :execresult
UPDATE `user`
SET role = ?
WHERE id = ? AND deleted_at IS NULL;

-- name: DeleteUser :execresult
UPDATE `user`
SET deleted_at = NOW()
WHERE id = ? AND deleted_at IS NULL;