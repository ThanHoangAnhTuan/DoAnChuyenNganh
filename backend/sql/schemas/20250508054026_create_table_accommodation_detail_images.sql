-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS `ecommerce_go_accommodation_detail_image` (
        `id` VARCHAR(36) NOT NULL COMMENT 'ID',
        `accommodation_detail_id` VARCHAR(36) NOT NULL COMMENT 'accommodation detail ID',
        `image` VARCHAR(255) NOT NULL COMMENT 'image',
        `created_at` BIGINT UNSIGNED NOT NULL COMMENT 'created at',
        `updated_at` BIGINT UNSIGNED NOT NULL COMMENT 'updated at',
        PRIMARY KEY (`id`),
        FOREIGN KEY (`accommodation_detail_id`) REFERENCES `ecommerce_go_accommodation_detail` (`id`) ON DELETE CASCADE
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'accommodation detail image table';

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `ecommerce_go_accommodation_detail_image`;

-- +goose StatementEnd