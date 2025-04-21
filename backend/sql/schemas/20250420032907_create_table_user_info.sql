-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS `ecommerce_go_user_info` (
        `id` VARCHAR(36) NOT NULL COMMENT 'ID',
        `account` VARCHAR(255) NOT NULL COMMENT 'account',
        `user_name` VARCHAR(255) NOT NULL DEFAULT "" COMMENT 'user name',
        `image` VARCHAR(255) NOT NULL DEFAULT "" COMMENT 'image',
        `status` TINYINT UNSIGNED NOT NULL COMMENT 'status: 0-locked, 1-actived, 2-not activated',
        `phone` VARCHAR(20) NOT NULL DEFAULT "" COMMENT 'phone',
        `gender` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT "gender: 0-male, 1-female",
        `birthday` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'birthday',
        `email` VARCHAR(255) NOT NULL DEFAULT "" COMMENT 'email',
        `is_authentication` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'authentication status: 0-not authenticated, 1-pending ,2-authenticated',
        `created_at` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'created at',
        `updated_at` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'updated at',
        PRIMARY KEY (`id`),
        UNIQUE KEY `unique_user_info_account` (`account`),
        UNIQUE KEY `unique_user_info_email` (`email`),
        UNIQUE KEY `unique_user_info_phone` (`phone`),
        FULLTEXT KEY `fulltext_user_info_account` (`account`),
        FULLTEXT KEY `fulltext_user_info_user_name` (`user_name`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'user info table';

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `ecommerce_go_user_info`;

-- +goose StatementEnd