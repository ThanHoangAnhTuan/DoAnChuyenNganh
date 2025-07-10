-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS `ecommerce_go_accommodation_room` (
        `id` VARCHAR(36) NOT NULL COMMENT 'ID',
        `accommodation_type` VARCHAR(36) NOT NULL COMMENT 'accommodation type',
        `name` VARCHAR(255) NOT NULL COMMENT 'accommodation type',
        `status` ENUM ('available', 'unavailable', 'occupied') NOT NULL COMMENT "status",
        `is_deleted` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'is deleted: 0 - not deleted; 1 - deleted',
        `created_at` BIGINT UNSIGNED NOT NULL COMMENT 'created at',
        `updated_at` BIGINT UNSIGNED NOT NULL COMMENT 'updated at',
        PRIMARY KEY (`id`),
        INDEX `idx_room_type_status_deleted` (`accommodation_type`, `status`, `is_deleted`),
        INDEX `idx_room_id_accommodation_type` (`id`, `accommodation_type`),
        FOREIGN KEY (`accommodation_type`) REFERENCES `ecommerce_go_accommodation_detail` (`id`) ON DELETE CASCADE,
        UNIQUE KEY `unique_accommodation_room_name_accommodation_type` (`name`, `accommodation_type`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'accommodation room table';

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `ecommerce_go_accommodation_room`;

-- +goose StatementEnd