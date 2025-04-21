-- name: CheckUserBaseExists :one
SELECT
    COUNT(*)
FROM
    `ecommerce_go_user_base`
WHERE
    `account` = ?;

-- name: GetUserBaseByAccount :one 
SELECT
    `id`,
    `account`,
    `password`
FROM
    `ecommerce_go_user_base`
WHERE
    `account` = ?
LIMIT
    1;

-- name: AddUserBase :exec
INSERT INTO
    `ecommerce_go_user_base` (
        `id`,
        `account`,
        `password`,
        `login_time`,
        `login_ip`,
        `logout_time`,
        `created_at`,
        `updated_at`
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?, ?);

-- name: LoginUserBase :exec
UPDATE `ecommerce_go_user_base`
SET
    `login_time` = ?,
    `login_ip` = ?
WHERE
    `account` = ?;

-- name: LogoutUserBase :exec
UPDATE `ecommerce_go_user_base`
SET
    `logout_time` = ?
WHERE
    `account` = ?;