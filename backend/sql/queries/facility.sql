-- name: CreateAccommodationFacility :exec
INSERT INTO
    `ecommerce_go_accommodation_facility` (`id`, `image`, `name`, `created_at`, `updated_at`)
VALUES
    (?, ?, ?, ?, ?);

-- name: GetAccommodationFacilityById :one
SELECT
    `id`,
    `image`,
    `name`
FROM
    `ecommerce_go_accommodation_facility`
WHERE
    `id` = ?;

-- name: GetAccommodationFacilityNames :many
SELECT
    `id`,
    `name`,
    `image`
FROM
    `ecommerce_go_accommodation_facility`;