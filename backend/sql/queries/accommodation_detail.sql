-- name: CreateAccommodationDetail :exec
INSERT INTO
    `ecommerce_go_accommodation_detail` (
        `id`,
        `accommodation_id`,
        `name`,
        `guests`,
        `beds`,
        `facilities`,
        `price`,
        `created_at`,
        `updated_at`
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetAccommodationDetail :one
SELECT
    `id`,
    `accommodation_id`,
    `name`,
    `guests`,
    `beds`,
    `facilities`,
    `price`
FROM
    `ecommerce_go_accommodation_detail`
WHERE
    `id` = ?
    and `accommodation_id` = ?
    and `is_deleted` = 0;

-- name: GetAccommodationDetails :many
SELECT
    `id`,
    `accommodation_id`,
    `name`,
    `guests`,
    `beds`,
    `discount_id`,
    `facilities`,
    `price`
FROM
    `ecommerce_go_accommodation_detail`
WHERE
    `accommodation_id` = ?
    and `is_deleted` = 0;

-- name: UpdateAccommodationDetail :exec
UPDATE `ecommerce_go_accommodation_detail`
SET
    `name` = ?,
    `guests` = ?,
    `beds` = ?,
    `facilities` = ?,
    `price` = ?,
    `updated_at` = ?
WHERE
    `id` = ?
    and `accommodation_id` = ?
    and `is_deleted` = 0;

-- name: DeleteAccommodationDetail :exec
UPDATE `ecommerce_go_accommodation_detail`
SET
    `is_deleted` = 1
WHERE
    `id` = ?;

-- name: CheckAccommodationDetailExists :one
SELECT
    EXISTS (
        SELECT
            1
        FROM
            `ecommerce_go_accommodation_detail`
        WHERE
            `id` = ?
    );

-- name: GetAccommodationDetailsByIDs :many
SELECT
    *
FROM
    `ecommerce_go_accommodation_detail`
WHERE
    `id` IN (sqlc.slice ('ids'))
    AND `accommodation_id` = ?;

-- name: CountAccommodationDetail :one
SELECT
    COUNT(*)
FROM
    `ecommerce_go_accommodation_detail`;

-- name: GetAccommodationDetailsWithPagination :many
SELECT
    `id`,
    `accommodation_id`,
    `name`,
    `guests`,
    `beds`,
    `discount_id`,
    `facilities`,
    `price`
FROM
    `ecommerce_go_accommodation_detail`
LIMIT
    ?
OFFSET
    ?;