-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS `ecommerce_go_accommodation_detail` (
        `id` VARCHAR(36) NOT NULL COMMENT 'ID',
        `accommodation_id` VARCHAR(36) NOT NULL COMMENT 'accommodation ID',
        `accommodation_type` VARCHAR(255) NOT NULL COMMENT 'accommodation type',
        `number_of_guests` TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT 'number of guests',
        `number_of_beds` TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT 'number of beds',
        `images` VARCHAR(255) NOT NULL DEFAULT "" COMMENT 'images',
        `facilities` JSON NOT NULL COMMENT 'facilities', -- coi lai
        `available_rooms` TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT 'available rooms',
        `price` DECIMAL(10, 2) NOT NULL COMMENT 'price',
        `discount_id` VARCHAR(36) COMMENT 'discount ID',
        `is_deleted` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'is deleted: 0 - not deleted; 1 - deleted',
        `created_at` BIGINT UNSIGNED NOT NULL COMMENT 'created at',
        `updated_at` BIGINT UNSIGNED NOT NULL COMMENT 'updated at',
        PRIMARY KEY (`id`),
        FOREIGN KEY (`accommodation_id`) REFERENCES `ecommerce_go_accommodation` (`id`) ON DELETE CASCADE,
        FOREIGN KEY (`discount_id`) REFERENCES `ecommerce_go_discount` (`id`) ON DELETE CASCADE,
        UNIQUE KEY `unique_accommodation_detail_id_discount_id` (`id`, `discount_id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'accommodation detail table';

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `ecommerce_go_accommodation_detail`;

-- +goose StatementEnd