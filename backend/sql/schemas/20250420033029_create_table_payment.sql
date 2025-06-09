-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS `ecommerce_go_payment` (
        `id` VARCHAR(36) NOT NULL COMMENT 'ID',
        `order_id` VARCHAR(36) NOT NULL COMMENT 'order ID',
        `payment_status` ENUM (
            'failed',
            'success',
            'refunded'
        ) NOT NULL COMMENT 'payment status',
        `payment_method` ENUM ('cash', 'card') NOT NULL COMMENT 'payment method',
        `total_price` DECIMAL(15, 2) NOT NULL COMMENT 'total price',
        `transaction_id` VARCHAR(100) COMMENT 'transaction id',
        `created_at` BIGINT UNSIGNED NOT NULL COMMENT 'created at',
        `updated_at` BIGINT UNSIGNED NOT NULL COMMENT 'updated at',
        PRIMARY KEY (`id`),
        FOREIGN KEY (`order_id`) REFERENCES `ecommerce_go_order` (`id`) ON DELETE CASCADE,
        UNIQUE KEY `unique_payment_order_id` (`order_id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'payment table';

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `ecommerce_go_payment`;

-- +goose StatementEnd