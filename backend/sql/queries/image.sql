-- name: SaveImage :exec
INSERT INTO
    `ecommerce_go_image` (
        `id`,
        `accommodation_detail_id`,
        `image`,
        `is_deleted`,
        `created_at`,
        `updated_at`
    )
VALUES
    (?, ?, ?, 0, ?, ?);

-- name: GetImages :many
SELECT
    `id`,
    `image`
FROM
    `ecommerce_go_image`
WHERE
    `accommodation_detail_id` = ? and `is_deleted` = 0;

-- name: DeleteImage :exec
UPDATE `ecommerce_go_image`
SET
    `is_deleted` = 1
WHERE
    `id` = ?;