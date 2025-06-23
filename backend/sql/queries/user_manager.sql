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

-- name: UpdateUserManagerLogin :exec
UPDATE `ecommerce_go_user_manager`
SET
    `login_time` = ?
WHERE
    `account` = ?;

-- name: GetUserManager :one
SELECT
    `id`,
    `account`,
    `user_name`,
    `password`
FROM
    `ecommerce_go_user_manager`
WHERE
    `account` = ?
    AND `is_deleted` = 0;

-- name: CheckUserManagerExistsByEmail :one
SELECT
    EXISTS (
        SELECT
            1
        FROM
            `ecommerce_go_user_manager`
        WHERE
            `account` = ?
            AND `is_deleted` = 0
    );

-- name: CheckUserManagerExistsByID :one
SELECT
    EXISTS (
        SELECT
            1
        FROM
            `ecommerce_go_user_manager`
        WHERE
            `id` = ?
            AND `is_deleted` = 0
    );

-- name: DeleteUserManager :exec
UPDATE `ecommerce_go_user_manager`
SET
    `is_deleted` = 1
WHERE
    `account` = ?;

-- name: IsAccommodationDetailBelongsToManager :one
SELECT
    EXISTS (
        SELECT
            1
        FROM
            `ecommerce_go_user_manager` m
            JOIN `ecommerce_go_accommodation` a ON m.id = a.manager_id
            JOIN `ecommerce_go_accommodation_detail` ad ON a.id = ad.accommodation_id
        WHERE
            m.id = ?
            AND ad.id = ?
    );

-- name: CountNumberOfManagers :one
SELECT
    COUNT(egum.id)
FROM
    `ecommerce_go_user_manager` egum;

-- name: GetManagers :many
SELECT
    `id`,
    `account`,
    `user_name`,
    `is_deleted`,
    `login_time`,
    `logout_time`,
    `created_at`,
    `updated_at`
FROM
    `ecommerce_go_user_manager`
WHERE
    `is_deleted` = 0
LIMIT
    ?
OFFSET
    ?;