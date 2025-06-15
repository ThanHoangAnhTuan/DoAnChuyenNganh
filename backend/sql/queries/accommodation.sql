-- name: CreateAccommodation :exec
INSERT INTO
    `ecommerce_go_accommodation` (
        `id`,
        `manager_id`,
        `country`,
        `name`,
        `city`,
        `district`,
        `description`,
        `facilities`,
        `gg_map`,
        `address`,
        `rating`,
        `rules`,
        `created_at`,
        `updated_at`
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetAccommodations :many
SELECT
    `id`,
    `manager_id`,
    `country`,
    `name`,
    `city`,
    `district`,
    `description`,
    `facilities`,
    `address`,
    `gg_map`,
    `rules`,
    `rating`
FROM
    `ecommerce_go_accommodation`
WHERE
    `is_deleted` = 0;

-- name: GetAccommodationsByManager :many
SELECT
    `id`,
    `manager_id`,
    `country`,
    `name`,
    `city`,
    `district`,
    `description`,
    `facilities`,
    `gg_map`,
    `address`,
    `rules`,
    `rating`
FROM
    `ecommerce_go_accommodation`
WHERE
    `is_deleted` = 0
    AND `manager_id` = ?;

-- name: GetAccommodationById :one
SELECT
    `id`,
    `manager_id`,
    `country`,
    `name`,
    `city`,
    `district`,
    `address`,
    `description`,
    `facilities`,
    `gg_map`,
    `rules`,
    `rating`
FROM
    `ecommerce_go_accommodation`
WHERE
    `id` = ?
    AND `is_deleted` = 0;

-- name: GetAccommodationsByCity :many
SELECT
    `id`,
    `country`,
    `name`,
    `city`,
    `district`,
    `address`,
    `gg_map`,
    `rating`
FROM
    `ecommerce_go_accommodation`
WHERE
    `city` = ?
    AND `is_deleted` = 0;

-- name: UpdateAccommodation :exec
UPDATE `ecommerce_go_accommodation`
SET
    `country` = ?,
    `name` = ?,
    `city` = ?,
    `district` = ?,
    `description` = ?,
    `facilities` = ?,
    `gg_map` = ?,
    `address` = ?,
    `rules` = ?,
    `updated_at` = ?
WHERE
    `id` = ?
    AND `is_deleted` = 0;

-- name: DeleteAccommodation :exec
UPDATE `ecommerce_go_accommodation`
SET
    `is_deleted` = 1,
    `updated_at` = ?
WHERE
    `id` = ? AND `is_deleted` = 0;

-- name: CheckAccommodationExists :one
SELECT
    EXISTS (
        SELECT
            1
        FROM
            `ecommerce_go_accommodation`
        WHERE
            `id` = ? AND `is_deleted` = 0
    );

-- name: CountAccommodation :one
SELECT
    COUNT(*)
FROM
    `ecommerce_go_accommodation`
WHERE
    `is_deleted` = 0;

-- name: CountAccommodationByCity :one
SELECT
    COUNT(*)
FROM
    `ecommerce_go_accommodation`
WHERE
    `city` = ?
    AND `is_deleted` = 0;

-- name: CountAccommodationByManager :one
SELECT
    COUNT(*)
FROM
    `ecommerce_go_accommodation`
WHERE
    `manager_id` = ?
    AND `is_deleted` = 0;

-- name: GetAccommodationsWithPagination :many
SELECT
    `id`,
    `manager_id`,
    `country`,
    `name`,
    `city`,
    `district`,
    `description`,
    `facilities`,
    `address`,
    `gg_map`,
    `rules`,
    `rating`
FROM
    `ecommerce_go_accommodation`
WHERE
    `is_deleted` = 0
LIMIT
    ?
OFFSET
    ?;

-- name: GetAccommodationsByCityWithPagination :many
SELECT
    `id`,
    `manager_id`,
    `country`,
    `name`,
    `city`,
    `district`,
    `description`,
    `facilities`,
    `address`,
    `gg_map`,
    `rules`,
    `rating`
FROM
    `ecommerce_go_accommodation`
WHERE
    `city` = ?
    AND `is_deleted` = 0
LIMIT
    ?
OFFSET
    ?;

-- name: GetAccommodationsByManagerWithPagination :many
SELECT
    `id`,
    `manager_id`,
    `country`,
    `name`,
    `city`,
    `district`,
    `description`,
    `facilities`,
    `gg_map`,
    `address`,
    `rules`,
    `rating`
FROM
    `ecommerce_go_accommodation`
WHERE
    `is_deleted` = 0
    AND `manager_id` = ?
LIMIT
    ?
OFFSET
    ?;

-- name: GetAccommodationNameById :one
SELECT
    `name`
FROM
    `ecommerce_go_accommodation`
WHERE
    `id` = ?
    AND `is_deleted` = 0;