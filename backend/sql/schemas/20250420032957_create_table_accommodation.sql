-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS `ecommerce_go_accommodation` (
        `id` VARCHAR(36) NOT NULL COMMENT 'ID',
        `manager_id` VARCHAR(36) NOT NULL COMMENT 'manager ID',
        `name` VARCHAR(255) NOT NULL COMMENT 'name',
        `country` VARCHAR(255) NOT NULL COMMENT 'country',
        `city` VARCHAR(255) NOT NULL COMMENT 'city',
        `district` VARCHAR(255) NOT NULL COMMENT 'district',
        `address` VARCHAR(255) NOT NULL COMMENT 'address',
        `description` TEXT NOT NULL COMMENT 'description',
        `facilities` JSON NOT NULL COMMENT 'facilities',
        `rating` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'rating',
        `gg_map` TEXT NOT NULL COMMENT 'google map address',
        -- `property_surroundings` JSON NOT NULL COMMENT 'property surroundings',
        `rules` JSON NOT NULL COMMENT 'rules',
        `is_verified` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'is verified: 0 - unverified, 1 - verified',
        `is_deleted` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'is deleted: 0 - not deleted; 1 - deleted',
        `created_at` BIGINT UNSIGNED NOT NULL COMMENT 'created at',
        `updated_at` BIGINT UNSIGNED NOT NULL COMMENT 'updated at',
        PRIMARY KEY (`id`),
        FOREIGN KEY (`manager_id`) REFERENCES `ecommerce_go_user_manager` (`id`) ON DELETE CASCADE
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'accommodation table';

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `ecommerce_go_accommodation`;

-- +goose StatementEnd