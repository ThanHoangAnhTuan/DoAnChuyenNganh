-- name: CreatePayment :exec
INSERT INTO
    `ecommerce_go_payment` (
        `id`,
        `order_id`,
        `payment_status`,
        `payment_method`,
        `total_price`,
        `transaction_id`,
        `created_at`,
        `updated_at`
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetPayment :one
SELECT
    `id`,
    `order_id`,
    `payment_status`,
    `payment_method`,
    `total_price`,
    `transaction_id`,
    `created_at`,
    `updated_at`
FROM
    `ecommerce_go_payment`
WHERE
    `order_id` = ?;

-- name: UpdatePaymentStatus :exec
UPDATE `ecommerce_go_payment`
SET
    `payment_status` = ?
WHERE
    `id` = ?
    and `order_id` = ?;

-- name: GetPaymentInfo :one
SELECT
    `id`,
    `order_id`,
    `payment_status`,
    `payment_method`,
    `total_price`,
    `transaction_id`
FROM
    `ecommerce_go_payment`
WHERE
    `order_id` = ?
    and `transaction_id` = ?
LIMIT
    1;