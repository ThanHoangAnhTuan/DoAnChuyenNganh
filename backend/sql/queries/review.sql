-- name: CreateReview :exec
INSERT INTO
    `ecommerce_go_review` (
        `id`,
        `user_id`,
        `accommodation_id`,
        `comment`,
        `rating`,
        `created_at`,
        `updated_at`
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?);

-- name: GetReviews :many
SELECT
    `id`,
    `comment`,
    `rating`,
    `manager_response`,
    `created_at`
FROM
    `ecommerce_go_review`
WHERE
    `accommodation_id` = ?;

-- name: UpdateReview :exec
UPDATE `ecommerce_go_review`
SET
    `comment` = ?,
    `rating` = ?,
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