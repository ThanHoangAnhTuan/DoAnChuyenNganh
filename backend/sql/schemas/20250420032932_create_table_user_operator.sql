-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS `ecommerce_go_user_operator` (
        `id` VARCHAR(36) NOT NULL COMMENT 'ID',
        `user_id` VARCHAR(36) NOT NULL COMMENT 'user ID',
        `user_type` ENUM ('admin', 'manager') NOT NULL,
        `created_at` BIGINT UNSIGNED NOT NULL COMMENT 'created at',
        `updated_at` BIGINT UNSIGNED NOT NULL COMMENT 'updated at',
        PRIMARY KEY (`id`),
        UNIQUE KEY `unique_user_operator_user_id_user_type` (`user_id`, `user_type`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'user operator table';

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `ecommerce_go_user_operator`;

-- +goose StatementEnd