CREATE TABLE IF NOT EXISTS `users` (
    `user_id` INT NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(50) NOT NULL UNIQUE,
    `password` VARCHAR(256) NOT NULL,
    `role` VARCHAR(50) NOT NULL DEFAULT 'user',
    `created_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`user_id`)
);

CREATE TABLE IF NOT EXISTS `figures` (
    `figure_id` INT NOT NULL AUTO_INCREMENT,
    `figure` VARCHAR(50) NOT NULL UNIQUE,
    `created_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`figure_id`)
);

CREATE TABLE IF NOT EXISTS `colors` (
    `color_id` INT NOT NULL AUTO_INCREMENT,
    `color` VARCHAR(50) NOT NULL UNIQUE,
    `created_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`color_id`)
);

CREATE TABLE IF NOT EXISTS `patterns` (
    `pattern_id` INT NOT NULL AUTO_INCREMENT,
    `pattern` VARCHAR(50) NOT NULL,
    `created_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`pattern_id`)
);

CREATE TABLE IF NOT EXISTS `categories` (
  `category_id` INT NOT NULL AUTO_INCREMENT,
  `size` VARCHAR(2) NOT NULL,
  `price` DOUBLE NOT NULL,
  `gender` VARCHAR(10) NOT NULL,
  `figure_id` INT NOT NULL,
  `color_id` INT NOT NULL,
  `pattern_id` INT NOT NULL,
  PRIMARY KEY (`category_id`, `pattern_id`, `color_id`, `figure_id`),
  INDEX `fk_categories_figures1_idx` (`figure_id`),
  INDEX `fk_categories_colors1_idx` (`color_id`),
  INDEX `fk_categories_patterns1_idx` (`pattern_id`),
  CONSTRAINT `fk_categories_figures1`
    FOREIGN KEY (`figure_id`)
    REFERENCES `figures` (`figure_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_categories_colors1`
    FOREIGN KEY (`color_id`)
    REFERENCES `colors` (`color_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_categories_patterns1`
    FOREIGN KEY (`pattern_id`)
    REFERENCES `patterns` (`pattern_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS `products` (
  `product_id` INT NOT NULL AUTO_INCREMENT,
  `gender` VARCHAR(2) NOT NULL,
  `created_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `category_id` INT NOT NULL,
  PRIMARY KEY (`product_id`, `category_id`),
  INDEX `fk_products_categories1_idx` (`category_id`),
  CONSTRAINT `fk_products_categories1`
    FOREIGN KEY (`category_id`)
    REFERENCES `categories` (`category_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS `carts` (
  `cart_id` INT NOT NULL AUTO_INCREMENT,
  `user_id` INT NOT NULL,
  `created_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `product_count` INT NOT NULL,
  `amount` DOUBLE NOT NULL,
  PRIMARY KEY (`cart_id`),
  INDEX `fk_carts_users1_idx` (`user_id`),
  CONSTRAINT `fk_carts_users1`
    FOREIGN KEY (`user_id`)
    REFERENCES `users` (`user_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS `orders` (
  `order_id` INT NOT NULL AUTO_INCREMENT,
  `address` VARCHAR(500) NOT NULL,
  `status` VARCHAR(45) NOT NULL DEFAULT 'IN_PROGRESS',
  `cart_id` INT NOT NULL,
  `created_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`order_id`),
  INDEX `fk_orders_carts1_idx` (`cart_id`),
  CONSTRAINT `fk_orders_carts1`
    FOREIGN KEY (`cart_id`)
    REFERENCES `carts` (`cart_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS `payments` (
  `payment_id` INT NOT NULL AUTO_INCREMENT,
  `amount` DOUBLE NOT NULL,
  `bank` VARCHAR(80) NOT NULL,
  `date` DATETIME NOT NULL,
  `orders_order_id` INT NOT NULL,
  `url` VARCHAR(500) NULL,
  `created_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`payment_id`),
  INDEX `fk_payments_orders1_idx` (`orders_order_id`),
  CONSTRAINT `fk_payments_orders1`
    FOREIGN KEY (`orders_order_id`)
    REFERENCES `orders` (`order_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS `products_has_carts` (
  `product_id` INT NOT NULL,
  `cart_id` INT NOT NULL,
  INDEX `fk_products_has_carts_carts1_idx` (`cart_id`),
  INDEX `fk_products_has_carts_products1_idx` (`product_id`),
  PRIMARY KEY (`product_id`, `cart_id`),
  CONSTRAINT `fk_products_has_carts_products1`
    FOREIGN KEY (`product_id`)
    REFERENCES `products` (`product_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_products_has_carts_carts1`
    FOREIGN KEY (`cart_id`)
    REFERENCES `carts` (`cart_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);