// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: 0002_ecommerce_users_profile.sql

package database

import (
	"context"
	"database/sql"
)

const updateUserProfile = `-- name: UpdateUserProfile :execresult
UPDATE ` + "`" + `user_profile` + "`" + `
SET
    first_name = ?,
    last_name = ?,
    phone_number = ?,
    address = ?
WHERE user_id = ?
`

type UpdateUserProfileParams struct {
	FirstName   sql.NullString
	LastName    sql.NullString
	PhoneNumber sql.NullString
	Address     sql.NullString
	UserID      int32
}

func (q *Queries) UpdateUserProfile(ctx context.Context, arg UpdateUserProfileParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateUserProfile,
		arg.FirstName,
		arg.LastName,
		arg.PhoneNumber,
		arg.Address,
		arg.UserID,
	)
}
