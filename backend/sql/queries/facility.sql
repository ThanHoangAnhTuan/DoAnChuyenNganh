-- name: CreateFacility :exec
INSERT INTO
    `ecommerce_go_facility` (`id`, `image`, `name`, `created_at`, `updated_at`)
VALUES
    (?, ?, ?, ?, ?);

-- name: GetFacilityById :one
SELECT
    `id`,
    `image`,
    `name`
FROM
    `ecommerce_go_facility`
WHERE
    `id` = ?;