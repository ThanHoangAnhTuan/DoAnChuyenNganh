-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS `ecommerce_go_accommodation_image` (
        `id` VARCHAR(36) NOT NULL COMMENT 'ID',
        `accommodation_id` VARCHAR(36) NOT NULL COMMENT 'accommodation ID',
        `image` VARCHAR(255) NOT NULL COMMENT 'image',
        `created_at` BIGINT UNSIGNED NOT NULL COMMENT 'created at',
        `updated_at` BIGINT UNSIGNED NOT NULL COMMENT 'updated at',
        PRIMARY KEY (`id`),
        FOREIGN KEY (`accommodation_id`) REFERENCES `ecommerce_go_accommodation` (`id`) ON DELETE CASCADE
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'accommodation image table';

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `ecommerce_go_accommodation_image`;

-- +goose StatementEnd