-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS `ecommerce_go_accommodation` (
        `id` VARCHAR(36) NOT NULL COMMENT 'ID',
        `manager_id` VARCHAR(36) NOT NULL COMMENT 'manager ID',
        `name` VARCHAR(255) NOT NULL COMMENT 'name',
        `city` VARCHAR(255) NOT NULL COMMENT 'city',
        `provine` VARCHAR(255) NOT NULL COMMENT 'city',
        `district` VARCHAR(255) NOT NULL COMMENT 'district',
        `images` VARCHAR(255) NOT NULL COMMENT 'images',
        `description` VARCHAR(255) NOT NULL COMMENT 'description',
        `facilities` JSON NOT NULL COMMENT 'facilities',
        `rating` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'rating',
        `gg_map` VARCHAR(255) NOT NULL COMMENT 'google map address',
        `property_surroundings` JSON NOT NULL COMMENT 'property surroundings',
        `rules` VARCHAR(255) NOT NULL COMMENT 'rules',
        `created_at` BIGINT UNSIGNED NOT NULL COMMENT 'created at',
        `updated_at` BIGINT UNSIGNED NOT NULL COMMENT 'updated at',
        PRIMARY KEY (`id`),
        FOREIGN KEY (`manager_id`) REFERENCES `ecommerce_go_user_manager` (`id`) ON DELETE CASCADE,
        UNIQUE KEY `unique_accommodation_manager_id` (`manager_id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'accommodation table';

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `ecommerce_go_accommodation`;

-- +goose StatementEnd