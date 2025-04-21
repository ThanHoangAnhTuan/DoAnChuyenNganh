-- name: CreateUserAdmin :exec
INSERT INTO
    `ecommerce_go_user_admin` (
        `id`,
        `account`,
        `password`,
        `created_at`,
        `updated_at`
    )
VALUES
    (?, ?, ?, ?, ?);

-- name: GetUserAdmin :one
SELECT
    `account`
FROM
    `ecommerce_go_user_admin`
WHERE
    `account` = ?
    AND `is_deleted` = 0;

-- name: DeleteUserAdmin :exec
UPDATE `ecommerce_go_user_admin`
SET
    `is_deleted` = 1
WHERE
    `account` = ?;