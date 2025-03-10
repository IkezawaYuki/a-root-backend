
-- +migrate Up
CREATE TABLE IF NOT EXISTS `customers` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `wordpress_url` VARCHAR(255) DEFAULT NULL,
    `start_date` DATETIME DEFAULT NULL,
    `facebook_token` VARCHAR(255) DEFAULT NULL,
    `instagram_token_status` INT NOT NULL DEFAULT 0,
    `instagram_account_id` VARCHAR(255) DEFAULT NULL,
    `instagram_account_name` VARCHAR(255) DEFAULT NULL,
    `stripe_subscription_id` VARCHAR(255) DEFAULT NULL,
    `stripe_customer_id` VARCHAR(255) DEFAULT NULL,
    `payment_type` INT NOT NULL DEFAULT 0,
    `payment_status` INT NOT NULL DEFAULT 0,
    `delete_hash_flag` INT NOT NULL DEFAULT 0,
    `dashboard_status` INT NOT NULL DEFAULT 0,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- +migrate Down
DROP TABLE IF EXISTS `customers`;