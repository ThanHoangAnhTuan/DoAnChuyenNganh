-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS `ecommerce_go_order_detail` (
        `id` VARCHAR(36) NOT NULL COMMENT 'ID',
        `order_id` VARCHAR(36) NOT NULL COMMENT 'order ID',
        `price` DECIMAL(15, 0) NOT NULL COMMENT 'price',
        `quantity` TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT 'quantity',
        `accommodation_detail_id` VARCHAR(36) NOT NULL COMMENT 'accommodation detail ID',
        `accommodation_room_id` VARCHAR(36) NOT NULL COMMENT 'accommodation room ID',
        `created_at` BIGINT UNSIGNED NOT NULL COMMENT 'created at',
        `updated_at` BIGINT UNSIGNED NOT NULL COMMENT 'updated at',
        PRIMARY KEY (`id`),
        INDEX `idx_order_detail_room` (`accommodation_room_id`),
        INDEX `idx_order_detail_detail` (`accommodation_detail_id`),
        INDEX `idx_order_detail_order` (`order_id`),
        FOREIGN KEY (`order_id`) REFERENCES `ecommerce_go_order` (`id`) ON DELETE CASCADE,
        FOREIGN KEY (`accommodation_detail_id`) REFERENCES `ecommerce_go_accommodation_detail` (`id`) ON DELETE CASCADE,
        FOREIGN KEY (`accommodation_room_id`) REFERENCES `ecommerce_go_accommodation_room` (`id`) ON DELETE CASCADE,
        UNIQUE KEY `unique_order_detail_order_id_accommodation_detail_id` (`order_id`, `accommodation_detail_id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'order detail table';

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `ecommerce_go_order_detail`;

-- +goose StatementEnd