package dao

import (
	"context"
	"strings"
)

type GetUsersForAdminParams struct {
	Role   NullUserRole
	Email  string
	Limit  int32
	Offset int32
}

type GetUsersForAdminRow struct {
	ID    int32
	Email string
	Role  NullUserRole
}

func (q *Queries) GetUsersForAdmin(ctx context.Context, arg GetUsersForAdminParams) ([]GetUsersForAdminRow, error) {
	var (
		getUsersForAdmin strings.Builder
		args             []interface{}
	)

	getUsersForAdmin.WriteString(`
		SELECT id, email, role
		FROM user
		WHERE deleted_at is NULL
	`)

	if arg.Role.Valid {
		getUsersForAdmin.WriteString(" AND role = ?")
		args = append(args, arg.Role)
	}

	if arg.Email != "" {
		getUsersForAdmin.WriteString(" AND email LIKE ?")
		args = append(args, arg.Email+"%")
	}

	getUsersForAdmin.WriteString(`
		ORDER BY created_at DESC
		LIMIT ?
		OFFSET ?
	`)
	args = append(args, arg.Limit, arg.Offset)

	rows, err := q.db.QueryContext(ctx, getUsersForAdmin.String(), args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUsersForAdminRow
	for rows.Next() {
		var i GetUsersForAdminRow
		if err := rows.Scan(&i.ID, &i.Email, &i.Role); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
