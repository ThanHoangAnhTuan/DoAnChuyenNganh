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
    `id`,
    `account`,
    `user_name`,
    `password`
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

-- name: UpdateUserAdminLogin :exec
UPDATE `ecommerce_go_user_admin`
SET
    `login_time` = ?
WHERE
    `account` = ?;

-- name: CheckUserAdminExistsById :one
SELECT
    EXISTS (
        SELECT
            1
        FROM
            `ecommerce_go_user_admin`
        WHERE
            `id` = ?
            AND `is_deleted` = 0
    );

-- name: CheckUserAdminExistsByEmail :one
SELECT
    EXISTS (
        SELECT
            1
        FROM
            `ecommerce_go_user_admin`
        WHERE
            `account` = ?
            AND `is_deleted` = 0
    );