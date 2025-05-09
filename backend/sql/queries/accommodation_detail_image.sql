-- name: UpdateAccommodationDetailImages :exec
INSERT INTO
    `ecommerce_go_accommodation_detail_image` (
        `id`,
        `accommodation_detail_id`,
        `image`,
        `created_at`,
        `updated_at`
    )
VALUES
    (?, ?, ?, ?, ?);

-- name: GetAccommodationDetailImages :many
SELECT
    `id`,
    `image`
FROM
    `ecommerce_go_accommodation_detail_image`
WHERE
    `accommodation_detail_id` = ?;

-- name: DeleteAccommodationDetailImage :exec
DELETE FROM `ecommerce_go_accommodation_detail_image`
WHERE
    `id` = ?;