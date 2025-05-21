-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS `ecommerce_go_accommodation_facility` (
        `id` VARCHAR(36) NOT NULL COMMENT 'ID',
        `image` VARCHAR(255) NOT NULL COMMENT 'image',
        `name` VARCHAR(255) NOT NULL COMMENT 'name',
        `created_at` BIGINT UNSIGNED NOT NULL COMMENT 'created at',
        `updated_at` BIGINT UNSIGNED NOT NULL COMMENT 'updated at',
        PRIMARY KEY (`id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'facility table';

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `ecommerce_go_accommodation_facility`;

-- +goose StatementEnd