-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS `ecommerce_go_user_verify` (
        `id` VARCHAR(36) NOT NULL COMMENT 'ID',
        `otp` VARCHAR(6) NOT NULL COMMENT 'OTP',
        `verify_key` VARCHAR(255) NOT NULL COMMENT 'email or SMS',
        `key_hash` VARCHAR(255) NOT NULL COMMENT 'hash email or SMS',
        `type` TINYINT UNSIGNED NOT NULL COMMENT 'type: 0 - SMS; 1 - email',
        `is_verified` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'is verified: 0 - unverified, 1 - verified',
        `is_deleted` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'is deleted: 0 - not deleted; 1 - deleted',
        `created_at` BIGINT UNSIGNED NOT NULL COMMENT 'created at',
        `updated_at` BIGINT UNSIGNED NOT NULL COMMENT 'updated at',
        PRIMARY KEY (`id`),
        UNIQUE KEY `unique_user_verify_verify_key` (`verify_key`),
        UNIQUE KEY `unique_user_verify_verify_key_otp` (`verify_key`, `otp`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'user verify table';

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `ecommerce_go_user_verify`;

-- +goose StatementEnd