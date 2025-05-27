-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS `ecommerce_go_review` (
        `id` VARCHAR(36) NOT NULL COMMENT 'ID',
        `user_id` VARCHAR(36) NOT NULL COMMENT 'user ID',
        `accommodation_id` VARCHAR(36) NOT NULL COMMENT 'accommodation ID',
        `comment` TEXT NOT NULL COMMENT 'comment',
        `rating` TINYINT UNSIGNED NOT NULL COMMENT 'rating',
        `manager_response` TEXT NOT NULL COMMENT 'manager response',
        `is_deleted` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'is deleted: 0 - not deleted; 1 - deleted',
        `created_at` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'created at',
        `updated_at` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'updated at',
        PRIMARY KEY (`id`),
        FOREIGN KEY (`user_id`) REFERENCES `ecommerce_go_user_base` (`id`) ON DELETE CASCADE,
        FOREIGN KEY (`accommodation_id`) REFERENCES `ecommerce_go_accommodation` (`id`) ON DELETE CASCADE,
        UNIQUE KEY `unique_review_user_id_accommodation_id` (`user_id`, `accommodation_id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'review table';

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `ecommerce_go_review`;

-- +goose StatementEnd