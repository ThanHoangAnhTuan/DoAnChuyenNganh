-- name: CreateUserOperator :exec
INSERT INTO
    `ecommerce_go_user_operator` (
        `id`,
        `user_id`,
        `user_type`,
        `created_at`,
        `updated_at`
    )
VALUES
    (?, ?, ?, ?, ?);