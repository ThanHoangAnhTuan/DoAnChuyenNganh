-- name: CreateUserInfo :exec
INSERT INTO
    `ecommerce_go_user_info` (
        `id`,
        `account`,
        `status`,
        `is_authentication`,
        `created_at`,
        `updated_at`
    )
VALUES
    (?, ?, ?, ?, ?, ?);

-- name: GetNameAndImageUserInfo :one
SELECT
    `user_name`,
    `image`
FROM
    `ecommerce_go_user_info`
WHERE
    `account` = ?
LIMIT
    1;

-- name: GetUserInfo :one
SELECT
    `id`,
    `account`,
    `user_name`,
    `image`,
    `status`,
    `phone`,
    `gender`,
    `birthday`,
    `email`,
    `is_authentication`
FROM
    `ecommerce_go_user_info`
WHERE
    `account` = ?
LIMIT
    1;

-- name: GetUserInfos :many
SELECT
    `id`,
    `account`,
    `user_name`,
    `image`,
    `status`,
    `phone`,
    `gender`,
    `birthday`,
    `email`,
    `is_authentication`
FROM
    `ecommerce_go_user_info`
WHERE
    `id` IN (?);

-- name: UpdateUserInfo :exec
UPDATE `ecommerce_go_user_info`
SET
    `user_name` = ?,
    `image` = ?,
    `phone` = ?,
    `gender` = ?,
    `birthday` = ?,
    `email` = ?,
    `updated_at` = ?
WHERE
    `id` = ?
    AND `is_authentication` = 1;

-- name: DeleteUserInfo :exec
DELETE FROM `ecommerce_go_user_info`
WHERE
    `id` = ?;