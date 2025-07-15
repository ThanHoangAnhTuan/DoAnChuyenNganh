-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS `ecommerce_go_order_room_booking` (
        `id` VARCHAR(36) NOT NULL COMMENT 'ID',
        `order_detail_id` VARCHAR(36) NOT NULL COMMENT 'order detail ID',
        `accommodation_room_id` VARCHAR(36) NOT NULL COMMENT 'accommodation room ID',
        `booking_status` ENUM (
            'reserved',
            'checked_in',
            'checked_out',
            'canceled'
        ) NOT NULL DEFAULT 'reserved' COMMENT 'booking status for this specific room',
        `created_at` BIGINT UNSIGNED NOT NULL COMMENT 'created at',
        `updated_at` BIGINT UNSIGNED NOT NULL COMMENT 'updated at',
		`is_deleted` TINYINT(1) NOT NULL DEFAULT 0 COMMENT 'soft delete flag',
        PRIMARY KEY (`id`),
        INDEX `idx_booking_order_detail` (`order_detail_id`),
        INDEX `idx_booking_room` (`accommodation_room_id`),
        INDEX `idx_booking_status` (`booking_status`),
        FOREIGN KEY (`order_detail_id`) REFERENCES `ecommerce_go_order_detail` (`id`) ON DELETE CASCADE,
        FOREIGN KEY (`accommodation_room_id`) REFERENCES `ecommerce_go_accommodation_room` (`id`) ON DELETE CASCADE
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'order room booking management table';

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `ecommerce_go_order_room_booking`;

-- +goose StatementEnd
