-- name: CreateUserManage :exec
INSERT INTO
    `ecommerce_go_user_manager` (
        `id`,
        `account`,
        `password`,
        `created_at`,
        `updated_at`
    )
VALUES
    (?, ?, ?, ?, ?);

-- name: GetUserManager :one
SELECT
    `account`
FROM
    `ecommerce_go_user_manager`
WHERE
    `account` = ?
    AND `is_deleted` = 0;

-- name: DeleteUserManager :exec
UPDATE `ecommerce_go_user_manager`
SET
    `is_deleted` = 1
WHERE
    `account` = ?;