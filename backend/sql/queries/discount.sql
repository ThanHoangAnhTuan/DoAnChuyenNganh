-- name: CreateDiscount :exec
INSERT INTO
    `ecommerce_go_discount` (
        `id`,
        `user_operator_id`,
        `name`,
        `description`,
        `discount_type`,
        `discount_value`,
        `is_active`,
        `created_at`,
        `updated_at`
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetDiscounts :many
SELECT
    `id`,
    `name`,
    `description`,
    `discount_type`,
    `discount_value`,
    `is_active`
FROM
    `ecommerce_go_discount`
WHERE
    `user_operator_id` = ?;

-- name: UpdateDiscount :exec
UPDATE `ecommerce_go_discount`
SET
    `name` = ?,
    `description` = ?,
    `discount_type` = ?,
    `discount_value` = ?,
    `is_active` = ?
WHERE
    `id` = ?
    and `user_operator_id` = ?;

-- name: UpdateDiscountStatus :exec
UPDATE `ecommerce_go_discount`
SET
    `is_active` = 1
WHERE
    `id` = ?
    and `user_operator_id` = ?;

-- name: DeleteDiscount :exec
UPDATE `ecommerce_go_discount`
SET
    `is_deleted` = 1
WHERE
    `id` = ?
    and `user_operator_id` = ?;