package dao

import (
	"context"
	"database/sql"
	"strings"
)

type GetProductsWithSearchForAdminParams struct {
	FromPrice int64
	ToPrice   int64
	Search    string
	Limit     int32
	Offset    int32
}

type GetProductsWithSearchForAdminRow struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       int64
	Quantity    int32
	ImageUrl    string
}

func (q *Queries) GetProductsWithSearchForAdmin(ctx context.Context, arg GetProductsWithSearchForAdminParams) ([]GetProductsWithSearchForAdminRow, error) {
	var (
		getProductsWithSearchForAdmin strings.Builder
		args                          []interface{}
	)

	getProductsWithSearchForAdmin.WriteString(`
		SELECT id, name, description, price, quantity, image_url
		FROM product
		WHERE 
			deleted_at IS NULL
	`)

	if arg.FromPrice > 0 || arg.ToPrice > 0 {
		getProductsWithSearchForAdmin.WriteString(" AND price BETWEEN ? AND ?")
		args = append(args, arg.FromPrice, arg.ToPrice)
	}

	if arg.Search != "" {
		getProductsWithSearchForAdmin.WriteString(" AND MATCH(name) AGAINST (? IN NATURAL LANGUAGE MODE)")
		args = append(args, arg.Search)
	}

	getProductsWithSearchForAdmin.WriteString(`
		ORDER BY created_at DESC
		LIMIT ?
		OFFSET ?
	`)
	args = append(args, arg.Limit, arg.Offset)

	rows, err := q.db.QueryContext(ctx, getProductsWithSearchForAdmin.String(), args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetProductsWithSearchForAdminRow
	for rows.Next() {
		var i GetProductsWithSearchForAdminRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.Quantity,
			&i.ImageUrl,
		); err != nil {
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
