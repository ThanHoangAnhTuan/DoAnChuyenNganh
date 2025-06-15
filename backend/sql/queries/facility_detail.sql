-- name: CreateAccommodationFacilityDetail :exec
INSERT INTO
    `ecommerce_go_accommodation_detail_facility` (`id`, `name`, `created_at`, `updated_at`)
VALUES
    (?, ?, ?, ?);

-- name: GetAccommodationFacilityDetailById :one
SELECT
    `id`,
    `name`
FROM
    `ecommerce_go_accommodation_detail_facility`
WHERE
    `id` = ?;

-- name: GetAccommodationFacilityDetail :many
SELECT
    `id`,
    `name`
FROM
    `ecommerce_go_accommodation_detail_facility`;

-- name: UpdateFacilityDetail :exec
UPDATE `ecommerce_go_accommodation_detail_facility`
SET
    `name` = sqlc.arg ("name"),
    `updated_at` = sqlc.arg ("updated_at")
WHERE
    `id` = sqlc.arg ("id");

-- name: DeleteFacilityDetail :exec
DELETE FROM `ecommerce_go_accommodation_detail_facility`
WHERE
    `id` = sqlc.arg ("id");