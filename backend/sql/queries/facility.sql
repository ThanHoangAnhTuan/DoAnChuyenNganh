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

-- name: UpdateFacility :exec
UPDATE `ecommerce_go_accommodation_facility`
SET
    `name` = sqlc.arg ("name"),
    `image` = sqlc.arg ("image"),
    `updated_at` = sqlc.arg ("updated_at")
WHERE
    `id` = sqlc.arg ("id");

-- name: DeleteFacility :exec
DELETE FROM `ecommerce_go_accommodation_facility`
WHERE
    `id` = sqlc.arg ("id");