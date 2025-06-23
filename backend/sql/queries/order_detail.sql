-- name: CreateOrderDetail :exec
INSERT INTO
    `ecommerce_go_order_detail` (
        `id`,
        `order_id`,
        `price`,
        `quantity`,
        `accommodation_detail_id`,
        `accommodation_room_id`,
        `created_at`,
        `updated_at`
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetOrderDetails :many
SELECT
    `id`,
    `order_id`,
    `price`,
    `quantity`,
    `accommodation_detail_id`,
    `created_at`,
    `updated_at`
FROM
    `ecommerce_go_order_detail`
WHERE
    `order_id` = ?;