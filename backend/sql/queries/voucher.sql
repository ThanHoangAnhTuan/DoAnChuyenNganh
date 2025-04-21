-- name: CreateVoucher :exec
INSERT INTO
    `ecommerce_go_voucher` (
        `id`,
        `user_operator_id`,
        `code`,
        `discount_type`,
        `discount_value`,
        `is_active`,
        `created_at`,
        `updated_at`
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetVouchers :many
SELECT
    `id`,
    `code`,
    `discount_type`,
    `discount_value`,
    `is_active`
FROM
    `ecommerce_go_voucher`
WHERE
    `user_operator_id` = ?;

-- name: UpdateVoucher :exec
UPDATE `ecommerce_go_voucher`
SET
    `code` = ?,
    `discount_type` = ?,
    `discount_value` = ?,
    `is_active` = ?
WHERE
    `id` = ?
    and `user_operator_id` = ?;

-- name: UpdateVoucherStatus :exec
UPDATE `ecommerce_go_voucher`
SET
    `is_active` = 1
WHERE
    `id` = ?
    and `user_operator_id` = ?;

-- name: DeleteVoucher :exec
UPDATE `ecommerce_go_voucher`
SET
    `is_deleted` = 1
WHERE
    `id` = ?
    and `user_operator_id` = ?;