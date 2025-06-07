-- name: CreateReview :exec
INSERT INTO
    `ecommerce_go_review` (
        `id`,
        `user_id`,
        `accommodation_id`,
        `title`,
        `comment`,
        `rating`,
        `created_at`,
        `updated_at`
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetReviews :many
SELECT
    `id`,
    `user_id`,
    `comment`,
    `rating`,
    `title`,
    `manager_response`,
    `created_at`
FROM
    `ecommerce_go_review`
WHERE
    `accommodation_id` = ?;

-- name: GetReviewsWithPagination :many
SELECT
    `id`,
    `user_id`,
    `comment`,
    `rating`,
    `title`,
    `manager_response`,
    `created_at`
FROM
    `ecommerce_go_review`
WHERE
    `accommodation_id` = ?
ORDER BY `created_at` DESC
LIMIT ? OFFSET ?;

-- name: UpdateReview :exec
UPDATE `ecommerce_go_review`
SET
    `comment` = ?,
    `rating` = ?,
    `title` = ?,
    `manager_response` = ?
WHERE
    `id` = ?
    and `user_id` = ?;

-- name: DeleteReview :exec
UPDATE `ecommerce_go_review`
SET
    `is_deleted` = 1
WHERE
    `id` = ?
    and `accommodation_id` = ?;

-- name: CountReviewsByAccommodation :one
SELECT COUNT(*)
FROM
    `ecommerce_go_review`
WHERE
    `accommodation_id` = ?