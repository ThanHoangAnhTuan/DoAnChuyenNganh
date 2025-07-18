// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: discount.sql

package database

import (
	"context"
	"database/sql"
)

const createDiscount = `-- name: CreateDiscount :exec
INSERT INTO
    ` + "`" + `ecommerce_go_discount` + "`" + ` (
        ` + "`" + `id` + "`" + `,
        ` + "`" + `user_operator_id` + "`" + `,
        ` + "`" + `name` + "`" + `,
        ` + "`" + `description` + "`" + `,
        ` + "`" + `discount_type` + "`" + `,
        ` + "`" + `discount_value` + "`" + `,
        ` + "`" + `is_active` + "`" + `,
        ` + "`" + `created_at` + "`" + `,
        ` + "`" + `updated_at` + "`" + `
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateDiscountParams struct {
	ID             string
	UserOperatorID string
	Name           string
	Description    sql.NullString
	DiscountType   EcommerceGoDiscountDiscountType
	DiscountValue  string
	IsActive       uint8
	CreatedAt      uint64
	UpdatedAt      uint64
}

func (q *Queries) CreateDiscount(ctx context.Context, arg CreateDiscountParams) error {
	_, err := q.db.ExecContext(ctx, createDiscount,
		arg.ID,
		arg.UserOperatorID,
		arg.Name,
		arg.Description,
		arg.DiscountType,
		arg.DiscountValue,
		arg.IsActive,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const deleteDiscount = `-- name: DeleteDiscount :exec
UPDATE ` + "`" + `ecommerce_go_discount` + "`" + `
SET
    ` + "`" + `is_deleted` + "`" + ` = 1
WHERE
    ` + "`" + `id` + "`" + ` = ?
    and ` + "`" + `user_operator_id` + "`" + ` = ?
`

type DeleteDiscountParams struct {
	ID             string
	UserOperatorID string
}

func (q *Queries) DeleteDiscount(ctx context.Context, arg DeleteDiscountParams) error {
	_, err := q.db.ExecContext(ctx, deleteDiscount, arg.ID, arg.UserOperatorID)
	return err
}

const getDiscounts = `-- name: GetDiscounts :many
SELECT
    ` + "`" + `id` + "`" + `,
    ` + "`" + `name` + "`" + `,
    ` + "`" + `description` + "`" + `,
    ` + "`" + `discount_type` + "`" + `,
    ` + "`" + `discount_value` + "`" + `,
    ` + "`" + `is_active` + "`" + `
FROM
    ` + "`" + `ecommerce_go_discount` + "`" + `
WHERE
    ` + "`" + `user_operator_id` + "`" + ` = ?
`

type GetDiscountsRow struct {
	ID            string
	Name          string
	Description   sql.NullString
	DiscountType  EcommerceGoDiscountDiscountType
	DiscountValue string
	IsActive      uint8
}

func (q *Queries) GetDiscounts(ctx context.Context, userOperatorID string) ([]GetDiscountsRow, error) {
	rows, err := q.db.QueryContext(ctx, getDiscounts, userOperatorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetDiscountsRow
	for rows.Next() {
		var i GetDiscountsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.DiscountType,
			&i.DiscountValue,
			&i.IsActive,
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

const updateDiscount = `-- name: UpdateDiscount :exec
UPDATE ` + "`" + `ecommerce_go_discount` + "`" + `
SET
    ` + "`" + `name` + "`" + ` = ?,
    ` + "`" + `description` + "`" + ` = ?,
    ` + "`" + `discount_type` + "`" + ` = ?,
    ` + "`" + `discount_value` + "`" + ` = ?,
    ` + "`" + `is_active` + "`" + ` = ?
WHERE
    ` + "`" + `id` + "`" + ` = ?
    and ` + "`" + `user_operator_id` + "`" + ` = ?
`

type UpdateDiscountParams struct {
	Name           string
	Description    sql.NullString
	DiscountType   EcommerceGoDiscountDiscountType
	DiscountValue  string
	IsActive       uint8
	ID             string
	UserOperatorID string
}

func (q *Queries) UpdateDiscount(ctx context.Context, arg UpdateDiscountParams) error {
	_, err := q.db.ExecContext(ctx, updateDiscount,
		arg.Name,
		arg.Description,
		arg.DiscountType,
		arg.DiscountValue,
		arg.IsActive,
		arg.ID,
		arg.UserOperatorID,
	)
	return err
}

const updateDiscountStatus = `-- name: UpdateDiscountStatus :exec
UPDATE ` + "`" + `ecommerce_go_discount` + "`" + `
SET
    ` + "`" + `is_active` + "`" + ` = 1
WHERE
    ` + "`" + `id` + "`" + ` = ?
    and ` + "`" + `user_operator_id` + "`" + ` = ?
`

type UpdateDiscountStatusParams struct {
	ID             string
	UserOperatorID string
}

func (q *Queries) UpdateDiscountStatus(ctx context.Context, arg UpdateDiscountStatusParams) error {
	_, err := q.db.ExecContext(ctx, updateDiscountStatus, arg.ID, arg.UserOperatorID)
	return err
}
