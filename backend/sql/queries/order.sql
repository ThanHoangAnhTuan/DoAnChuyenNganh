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
            `user_id` = ?
            and `order_id_external` = ?
    );

-- name: UpdateOrderStatus :exec
UPDATE `ecommerce_go_order`
SET
    `order_status` = ?,
    `updated_at` = ?
WHERE
    `order_id_external` = ?;

-- name: GetOrderIdAndUserIdByOrderIdExternal :one
SELECT
    `id`,
    `user_id`
FROM
    `ecommerce_go_order`
WHERE
    `order_id_external` = ?;

-- name: GetOrderInfoByOrderIDExternal :one
SELECT
    `id`,
    `user_id`,
    `order_id_external`,
    `final_total`,
    `order_status`,
    `checkin_date`,
    `checkout_date`,
    `created_at`
FROM
    `ecommerce_go_order`
WHERE
    `order_id_external` = ?
LIMIT
    1;

-- name: GetOrdersByUser :many
SELECT
    ego.id AS order_id,
    ego.final_total,
    ego.order_status,
    ego.checkin_date,
    ego.checkout_date,
    ega.name AS accommodation_name,
    ega.id AS accommodation_id
FROM
    `ecommerce_go_order` ego
    JOIN `ecommerce_go_accommodation` ega ON ego.accommodation_id = ega.id
WHERE
    ego.user_id = ?
ORDER BY
    FIELD (
        ego.order_status,
        'payment_success',
        'checked_in',
        'completed',
        'pending_payment',
        'canceled',
        'refunded',
        'payment_failed'
    ),
    ego.created_at ASC;

-- name: GetOrderDetailsByUser :many
SELECT
    egod.accommodation_detail_id AS `accommodation_detail_id`,
    egod.price AS `price`,
    egad.name AS `accommodation_detail_name`,
    egad.guests AS `guests`
FROM
    `ecommerce_go_order_detail` egod
    JOIN `ecommerce_go_accommodation_detail` egad ON egod.accommodation_detail_id = egad.id
WHERE
    egod.order_id = sqlc.arg ('order_id')
ORDER BY
    egod.created_at ASC;

-- name: GetOrdersByManager :many
SELECT
    ego.id AS order_id,
    ego.final_total,
    ego.order_status,
    ego.checkin_date,
    ego.checkout_date,
    egui.account AS email,
    egui.user_name AS username,
    egui.phone AS phone,
    ega.name AS accommodation_name,
    ega.id AS accommodation_id
FROM
    `ecommerce_go_order` ego
    JOIN `ecommerce_go_accommodation` ega ON ego.accommodation_id = ega.id
    JOIN `ecommerce_go_user_manager` egum ON egum.id = ega.manager_id
    JOIN `ecommerce_go_user_info` egui ON egui.id = ego.user_id
WHERE
    egum.id = sqlc.arg ('manager_id')
ORDER BY
    FIELD (
        ego.order_status,
        'payment_success',
        'checked_in',
        'completed',
        'pending_payment',
        'canceled',
        'refunded',
        'payment_failed'
    ),
    ego.created_at ASC;

-- name: GetOrderDetailsByManager :many
SELECT
    egod.accommodation_detail_id AS `accommodation_detail_id`,
    egod.price AS `price`,
    egad.name AS `accommodation_detail_name`
FROM
    `ecommerce_go_order_detail` egod
    JOIN `ecommerce_go_accommodation_detail` egad ON egod.accommodation_detail_id = egad.id
WHERE
    egod.order_id = sqlc.arg ('order_id')
ORDER BY
    egod.created_at ASC;

-- name: CheckOrderExists :one
SELECT
    EXISTS (
        SELECT
            1
        FROM
            `ecommerce_go_order`
        WHERE
            `id` = ?
    );

-- name: UpdateOrderStatusByID :exec
UPDATE `ecommerce_go_order`
SET
    `order_status` = ?,
    `updated_at` = ?
WHERE
    `id` = ?;

