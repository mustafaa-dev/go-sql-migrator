-- Query 1: query
CREATE TABLE IF NOT EXISTS `admins` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `uuid` VARCHAR(30) NULL,
  `name` VARCHAR(255) NOT NULL,
  `email` VARCHAR(255) NULL,
  `email_verified_at` TIMESTAMP NULL,
  `phone_dial_code` VARCHAR(10) NOT NULL,
  `phone_number` VARCHAR(45) NOT NULL,
  `phone_verified_at` TIMESTAMP NULL,
  `password` VARCHAR(255) NOT NULL,
  `role` SMALLINT NOT NULL,
  `image` VARCHAR(255) NULL,
  `status` ENUM('active', 'suspended') NULL DEFAULT 'active',
  `remember_token` VARCHAR(100) NULL,
  `created_at` TIMESTAMP NULL,
  `updated_at` TIMESTAMP NULL,
  `deleted_at` TIMESTAMP NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `uuid_UNIQUE` (`uuid` ASC) VISIBLE,
  UNIQUE INDEX `email_UNIQUE` (`email` ASC) VISIBLE,
  UNIQUE INDEX `phone_dial_code_UNIQUE` (`phone_dial_code` ASC, `phone_number` ASC) VISIBLE)
ENGINE = InnoDB;

-- Query 2: query
CREATE TABLE IF NOT EXISTS `authorities` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `admin_id` BIGINT NOT NULL,
  `authority` VARCHAR(255),
  PRIMARY KEY (`id`),
  INDEX `fk_authorities_admins_idx` (`admin_id` ASC) VISIBLE,
  CONSTRAINT `fk_authorities_admins`
    FOREIGN KEY (`admin_id`)
    REFERENCES `admins` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;