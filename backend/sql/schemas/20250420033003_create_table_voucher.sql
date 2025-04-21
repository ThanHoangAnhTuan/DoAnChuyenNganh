-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS `ecommerce_go_voucher` (
        `id` VARCHAR(36) NOT NULL COMMENT 'ID',
        `user_operator_id` VARCHAR(36) NOT NULL COMMENT 'user operator ID',
        `code` VARCHAR(36) NOT NULL COMMENT 'code',
        `discount_type` ENUM ('fixed', 'percentage') NOT NULL COMMENT 'discount type',
        `discount_value` DECIMAL(10, 2) NOT NULL COMMENT 'discount value',
        `is_deleted` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'is deleted: 0 - not deleted; 1 - deleted',
        `is_active` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '0 - inactivated; 1 - activated',
        `created_at` BIGINT UNSIGNED NOT NULL COMMENT 'created at',
        `updated_at` BIGINT UNSIGNED NOT NULL COMMENT 'updated at',
        PRIMARY KEY (`id`),
        FOREIGN KEY (`user_operator_id`) REFERENCES `ecommerce_go_user_operator` (`id`) ON DELETE CASCADE
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'voucher table';

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `ecommerce_go_voucher`;

-- +goose StatementEnd