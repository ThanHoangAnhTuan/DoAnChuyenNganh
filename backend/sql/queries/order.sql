-- name: CreateOrder :exec
INSERT INTO
    `ecommerce_go_order` (
        `id`,
        `user_id`,
        `total_price`,
        `order_status`,
        `voucher_id`,
        `checkin_date`,
        `checkout_date`,
        `created_at`,
        `updated_at`
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetOrdersByUser :many
SELECT
    `id`,
    `total_price`,
    `order_status`,
    `voucher_id`,
    `checkin_date`,
    `checkout_date`,
    `created_at`,
    `updated_at`
FROM
    `ecommerce_go_order`
WHERE
    `user_id` = ?;

-- name: UpdateOrderStatus :exec
UPDATE `ecommerce_go_order`
SET
    `order_status` = ?,
    `updated_at` = ?
WHERE
    `id` = ?;