// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: review.sql

package database

import (
	"context"
	"database/sql"
)

const countReviewsByAccommodation = `-- name: CountReviewsByAccommodation :one
SELECT COUNT(*)
FROM
    ` + "`" + `ecommerce_go_review` + "`" + `
WHERE
    ` + "`" + `accommodation_id` + "`" + ` = ?
`

func (q *Queries) CountReviewsByAccommodation(ctx context.Context, accommodationID string) (int64, error) {
	row := q.db.QueryRowContext(ctx, countReviewsByAccommodation, accommodationID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createReview = `-- name: CreateReview :exec
INSERT INTO
    ` + "`" + `ecommerce_go_review` + "`" + ` (
        ` + "`" + `id` + "`" + `,
        ` + "`" + `user_id` + "`" + `,
        ` + "`" + `accommodation_id` + "`" + `,
        ` + "`" + `title` + "`" + `,
        ` + "`" + `comment` + "`" + `,
        ` + "`" + `rating` + "`" + `,
        ` + "`" + `created_at` + "`" + `,
        ` + "`" + `updated_at` + "`" + `
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateReviewParams struct {
	ID              string
	UserID          string
	AccommodationID string
	Title           string
	Comment         string
	Rating          uint8
	CreatedAt       uint64
	UpdatedAt       uint64
}

func (q *Queries) CreateReview(ctx context.Context, arg CreateReviewParams) error {
	_, err := q.db.ExecContext(ctx, createReview,
		arg.ID,
		arg.UserID,
		arg.AccommodationID,
		arg.Title,
		arg.Comment,
		arg.Rating,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const deleteReview = `-- name: DeleteReview :exec
UPDATE ` + "`" + `ecommerce_go_review` + "`" + `
SET
    ` + "`" + `is_deleted` + "`" + ` = 1
WHERE
    ` + "`" + `id` + "`" + ` = ?
    and ` + "`" + `accommodation_id` + "`" + ` = ?
`

type DeleteReviewParams struct {
	ID              string
	AccommodationID string
}

func (q *Queries) DeleteReview(ctx context.Context, arg DeleteReviewParams) error {
	_, err := q.db.ExecContext(ctx, deleteReview, arg.ID, arg.AccommodationID)
	return err
}

const getReviews = `-- name: GetReviews :many
SELECT
    ` + "`" + `id` + "`" + `,
    ` + "`" + `user_id` + "`" + `,
    ` + "`" + `comment` + "`" + `,
    ` + "`" + `rating` + "`" + `,
    ` + "`" + `title` + "`" + `,
    ` + "`" + `manager_response` + "`" + `,
    ` + "`" + `created_at` + "`" + `
FROM
    ` + "`" + `ecommerce_go_review` + "`" + `
WHERE
    ` + "`" + `accommodation_id` + "`" + ` = ?
`

type GetReviewsRow struct {
	ID              string
	UserID          string
	Comment         string
	Rating          uint8
	Title           string
	ManagerResponse sql.NullString
	CreatedAt       uint64
}

func (q *Queries) GetReviews(ctx context.Context, accommodationID string) ([]GetReviewsRow, error) {
	rows, err := q.db.QueryContext(ctx, getReviews, accommodationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetReviewsRow
	for rows.Next() {
		var i GetReviewsRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Comment,
			&i.Rating,
			&i.Title,
			&i.ManagerResponse,
			&i.CreatedAt,
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

const getReviewsWithPagination = `-- name: GetReviewsWithPagination :many
SELECT
    ` + "`" + `id` + "`" + `,
    ` + "`" + `user_id` + "`" + `,
    ` + "`" + `comment` + "`" + `,
    ` + "`" + `rating` + "`" + `,
    ` + "`" + `title` + "`" + `,
    ` + "`" + `manager_response` + "`" + `,
    ` + "`" + `created_at` + "`" + `
FROM
    ` + "`" + `ecommerce_go_review` + "`" + `
WHERE
    ` + "`" + `accommodation_id` + "`" + ` = ?
ORDER BY ` + "`" + `created_at` + "`" + ` DESC
LIMIT ? OFFSET ?
`

type GetReviewsWithPaginationParams struct {
	AccommodationID string
	Limit           int32
	Offset          int32
}

type GetReviewsWithPaginationRow struct {
	ID              string
	UserID          string
	Comment         string
	Rating          uint8
	Title           string
	ManagerResponse sql.NullString
	CreatedAt       uint64
}

func (q *Queries) GetReviewsWithPagination(ctx context.Context, arg GetReviewsWithPaginationParams) ([]GetReviewsWithPaginationRow, error) {
	rows, err := q.db.QueryContext(ctx, getReviewsWithPagination, arg.AccommodationID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetReviewsWithPaginationRow
	for rows.Next() {
		var i GetReviewsWithPaginationRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Comment,
			&i.Rating,
			&i.Title,
			&i.ManagerResponse,
			&i.CreatedAt,
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

const updateReview = `-- name: UpdateReview :exec
UPDATE ` + "`" + `ecommerce_go_review` + "`" + `
SET
    ` + "`" + `comment` + "`" + ` = ?,
    ` + "`" + `rating` + "`" + ` = ?,
    ` + "`" + `title` + "`" + ` = ?,
    ` + "`" + `manager_response` + "`" + ` = ?
WHERE
    ` + "`" + `id` + "`" + ` = ?
    and ` + "`" + `user_id` + "`" + ` = ?
`

type UpdateReviewParams struct {
	Comment         string
	Rating          uint8
	Title           string
	ManagerResponse sql.NullString
	ID              string
	UserID          string
}

func (q *Queries) UpdateReview(ctx context.Context, arg UpdateReviewParams) error {
	_, err := q.db.ExecContext(ctx, updateReview,
		arg.Comment,
		arg.Rating,
		arg.Title,
		arg.ManagerResponse,
		arg.ID,
		arg.UserID,
	)
	return err
}
