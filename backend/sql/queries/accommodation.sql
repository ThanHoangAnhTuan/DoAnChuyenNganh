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
    `is_deleted` = 0 AND `is_verified` = 1;

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
    AND `is_deleted` = 0 AND `is_verified` = 1;


-- name: GetAccommodationByIdNoVerify :one
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
    `rating`,
    `is_deleted`,
    `is_verified`
FROM
    `ecommerce_go_accommodation`
WHERE
    `id` = ?
    AND `is_deleted` = 0;

-- name: GetAccommodationByIdByAdmin :one
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
    `rating`,
    `is_deleted`,
    `is_verified`
FROM
    `ecommerce_go_accommodation`
WHERE
    `id` = ?;


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
    AND `is_deleted` = 0 AND `is_verified` = 1;

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
    `id` = ?
    AND `is_deleted` = 0;

-- name: RestoreAccommodation :exec
UPDATE `ecommerce_go_accommodation`
SET
    `is_deleted` = 0,
    `updated_at` = ?
WHERE
    `id` = ?
    AND `is_deleted` = 1;

-- name: CheckAccommodationExists :one
SELECT
    EXISTS (
        SELECT
            1
        FROM
            `ecommerce_go_accommodation`
        WHERE
            `id` = ?
            AND `is_deleted` = 0
    );

-- name: CheckAccommodationExistsByAdmin :one
SELECT
    EXISTS (
        SELECT
            1
        FROM
            `ecommerce_go_accommodation`
        WHERE
            `id` = ?
    );

-- name: CountAccommodation :one
SELECT
    COUNT(*)
FROM
    `ecommerce_go_accommodation`
WHERE
    `is_deleted` = 0 AND `is_verified` = 1;

-- name: CountAccommodationOfManager :one
SELECT
    COUNT(*)
FROM
    `ecommerce_go_accommodation`
WHERE
    `manager_id` = sqlc.arg ("manager_id");

-- name: CountAccommodationByCity :one
SELECT
    COUNT(*)
FROM
    `ecommerce_go_accommodation`
WHERE
    `city` = ?
    AND `is_deleted` = 0 AND `is_verified` = 1;

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
    `rating`,
    `is_deleted`,
    `is_verified`
FROM
    `ecommerce_go_accommodation`
WHERE
    `is_deleted` = 0 AND `is_verified` = 1
LIMIT
    ?
OFFSET
    ?;

-- name: GetAccommodationsOfManagerWithPagination :many
SELECT
    *
FROM
    `ecommerce_go_accommodation`
WHERE
    `manager_id` = ?
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
    `rating`,
    `is_deleted`,
    `is_verified`
FROM
    `ecommerce_go_accommodation`
WHERE
    `city` = ?
    AND `is_deleted` = 0 AND `is_verified` = 1
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
    `rating`,
    `is_deleted`,
    `is_verified`
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

-- name: UpdateStatusAccommodation :exec
UPDATE `ecommerce_go_accommodation`
SET
    `is_verified` = sqlc.arg ("is_verified")
WHERE
    `id` = sqlc.arg ("id");