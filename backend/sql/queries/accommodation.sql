-- name: CreateAccommodation :exec
INSERT INTO
    `ecommerce_go_accommodation` (
        `id`,
        `manager_id`,
        `country`,
        `name`,
        `city`,
        `district`,
        `image`,
        `description`,
        `facilities`,
        `gg_map`,
        `property_surroundings`,
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
    `image`,
    `description`,
    `facilities`,
    `gg_map`,
    `property_surroundings`,
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
    `image`,
    `description`,
    `facilities`,
    `gg_map`,
    `property_surroundings`,
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
    `image`,
    `description`,
    `facilities`,
    `gg_map`,
    `property_surroundings`,
    `rules`,
    `rating`
FROM
    `ecommerce_go_accommodation`
WHERE
    `id` = ?
    AND `is_deleted` = 0;

-- name: UpdateAccommodation :exec
UPDATE `ecommerce_go_accommodation`
SET
    `country` = ?,
    `name` = ?,
    `city` = ?,
    `district` = ?,
    `image` = ?,
    `description` = ?,
    `facilities` = ?,
    `gg_map` = ?,
    `property_surroundings` = ?,
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
    `id` = ?;

-- name: CheckAccommodationExists :one
SELECT
    EXISTS (
        SELECT
            1
        FROM
            `ecommerce_go_accommodation`
        WHERE
            `id` = ?
    );