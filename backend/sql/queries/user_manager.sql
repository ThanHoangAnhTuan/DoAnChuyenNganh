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

-- name: CheckUserManagerExistsByID :one
SELECT
    COUNT(*)
FROM
    `ecommerce_go_user_manager`
WHERE
    `id` = ?
    AND `is_deleted` = 0;

-- name: DeleteUserManager :exec
UPDATE `ecommerce_go_user_manager`
SET
    `is_deleted` = 1
WHERE
    `account` = ?;