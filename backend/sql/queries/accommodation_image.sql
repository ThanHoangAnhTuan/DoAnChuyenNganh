-- name: UpdateAccommodationImages :exec
INSERT INTO
    `ecommerce_go_accommodation_image` (
        `id`,
        `accommodation_id`,
        `image`,
        `created_at`,
        `updated_at`
    )
VALUES
    (?, ?, ?, ?, ?);

-- name: GetAccommodationImages :many
SELECT
    `id`,
    `image`
FROM
    `ecommerce_go_accommodation_image`
WHERE
    `accommodation_id` = ?;

-- name: GetAccommodationImage :one
SELECT
    `image`
FROM
    `ecommerce_go_accommodation_image`
WHERE
    `accommodation_id` = ?
LIMIT
    1;

-- name: DeleteAccommodationImage :exec
DELETE FROM `ecommerce_go_accommodation_image`
WHERE
    `image` = ?;