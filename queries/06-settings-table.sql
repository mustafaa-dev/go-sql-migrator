
CREATE TABLE settings (
                          id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
                          `key` VARCHAR(255) NOT NULL UNIQUE,
                          `value` TEXT,
                          `group` VARCHAR(255) DEFAULT NULL,
                          `type` VARCHAR(50) DEFAULT 'string',
                          `description` TEXT DEFAULT NULL,
                          `created_at` TIMESTAMP NULL DEFAULT NULL,
                          `updated_at` TIMESTAMP NULL DEFAULT NULL
);
