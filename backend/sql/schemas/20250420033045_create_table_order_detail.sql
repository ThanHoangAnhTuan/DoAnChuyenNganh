-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS `ecommerce_go_order_detail` (
        `id` VARCHAR(36) NOT NULL COMMENT 'ID',
        `order_id` VARCHAR(36) NOT NULL COMMENT 'order ID',
        `price` INT UNSIGNED NOT NULL COMMENT 'price',
        `accommodation_detail_id` VARCHAR(36) NOT NULL COMMENT 'accommodation detail ID',
        `created_at` BIGINT UNSIGNED NOT NULL COMMENT 'created at',
        `updated_at` BIGINT UNSIGNED NOT NULL COMMENT 'updated at',
        PRIMARY KEY (`id`),
        FOREIGN KEY (`order_id`) REFERENCES `ecommerce_go_order` (`id`) ON DELETE CASCADE,
        FOREIGN KEY (`accommodation_detail_id`) REFERENCES `ecommerce_go_accommodation_detail` (`id`) ON DELETE CASCADE,
        UNIQUE KEY `unique_order_detail_order_id_accommodation_detail_id` (`order_id`, `accommodation_detail_id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'order detail table';

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `ecommerce_go_order_detail`;

-- +goose StatementEnd