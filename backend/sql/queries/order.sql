-- name: CreateOrder :exec
INSERT INTO
    `ecommerce_go_order` (
        `id`,
        `user_id`,
        `order_id_external`,
        `final_total`,
        `order_status`,
        `accommodation_id`,
        `voucher_id`,
        `checkin_date`,
        `checkout_date`,
        `created_at`,
        `updated_at`
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: CheckUserBookedOrder :one
SELECT
    EXISTS (
        SELECT
            1
        FROM
            `ecommerce_go_order`
        WHERE
            `user_id` = ? and `order_id_external` = ? and `order_status` = `completed`
    );

-- name: GetOrdersByUser :many
SELECT
    `id`,
    `final_total`,
    `order_status`,
    -- `voucher_id`,
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
    `order_id_external` = ?;

-- name: GetOrderIdByOrderIdExternal :one
SELECT
    `id`
FROM
    `ecommerce_go_order`
WHERE
    `order_id_external` = ?;