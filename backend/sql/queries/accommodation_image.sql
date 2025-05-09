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

-- name: DeleteAccommodationImage :exec
DELETE FROM `ecommerce_go_accommodation_image`
WHERE
    `image` = ?;