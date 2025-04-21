-- name: CreateUserVerify :exec
INSERT INTO
    `ecommerce_go_user_verify` (
        `id`,
        `otp`,
        `verify_key`,
        `key_hash`,
        `type`,
        `is_verified`,
        `is_deleted`,
        `created_at`,
        `updated_at`
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetUserUnverify :one
SELECT
    `otp`,
    `key_hash`,
    `verify_key`,
    `is_verified`,
    `id`
FROM
    `ecommerce_go_user_verify`
WHERE
    `key_hash` = ?
    AND `is_verified` = 0;

-- name: GetUserVerified :one
SELECT
    `otp`,
    `key_hash`,
    `verify_key`,
    `is_verified`,
    `id`
FROM
    `ecommerce_go_user_verify`
WHERE
    `key_hash` = ?
    AND `is_verified` = 1;

-- name: UpdateUserVerifyStatus :exec
UPDATE `ecommerce_go_user_verify`
SET
    `is_verified` = 1,
    `is_deleted` = 1,
    `updated_at` = ?
WHERE
    `key_hash` = ?;