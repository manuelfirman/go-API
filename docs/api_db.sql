-- DDL
DROP DATABASE IF EXISTS `go_api_db`;

CREATE DATABASE `go_api_db`;

USE `go_api_db`;

-- table `product_types`
CREATE TABLE `localities` (
    `id` int NOT NULL,
    `locality_name` varchar(50) NOT NULL,
    `province_name` varchar(50) NOT NULL,
    `country_name` varchar(50) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = UTF8MB4;

-- table `carries`
CREATE TABLE `carries` (
    `id` int NOT NULL AUTO_INCREMENT,
    `cid` int NOT NULL,
    `company_name` varchar(255) NOT NULL,
    `address` varchar(255) NOT NULL,
    `telephone` varchar(15) NOT NULL,
    `locality_id` int NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_carriers_cid` (`cid`),
    CONSTRAINT `fk_carriers_locality_id` FOREIGN KEY (`locality_id`) REFERENCES `localities` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = UTF8MB4;

-- table `sellers`
CREATE TABLE `sellers` (
    `id` int NOT NULL AUTO_INCREMENT,
    `cid` int NOT NULL,
    `company_name` varchar(255) NOT NULL,
    `address` varchar(255) NOT NULL,
    `telephone` varchar(15) NOT NULL,
    `locality_id` int NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_sellers_cid` (`cid`),
    CONSTRAINT `fk_sellers_locality_id` FOREIGN KEY (`locality_id`) REFERENCES `localities` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = UTF8MB4;

-- table `warehouses`
CREATE TABLE `warehouses` (
    `id` int NOT NULL AUTO_INCREMENT,
    `warehouse_code` varchar(25) NOT NULL,
    `address` varchar(255) NOT NULL,
    `telephone` varchar(15) NOT NULL,
    `minimum_capacity` int NOT NULL,
    `minimum_temperature` float NOT NULL,
    `locality_id` int NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_warehouses_warehouse_code` (`warehouse_code`),
    CONSTRAINT `fk_warehouses_locality_id` FOREIGN KEY (`locality_id`) REFERENCES `localities` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = UTF8MB4;

-- table `sections`
CREATE TABLE `sections` (
    `id` int NOT NULL AUTO_INCREMENT,
    `section_number` int NOT NULL,
    `current_temperature` float NOT NULL,
    `minimum_temperature` float NOT NULL,
    `current_capacity` int NOT NULL,
    `minimum_capacity` int NOT NULL,
    `maximum_capacity` int NOT NULL,
    `warehouse_id` int NOT NULL,
    `product_type_id` int NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_sections_section_number` (`section_number`),
    CONSTRAINT `fk_sections_warehouse_id` FOREIGN KEY (`warehouse_id`) REFERENCES `warehouses` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = UTF8MB4;

-- table `products`
CREATE TABLE `products` (
    `id` int NOT NULL AUTO_INCREMENT,
    `product_code` varchar(25) NOT NULL,
    `description` text NOT NULL,
    `height` float NOT NULL,
    `length` float NOT NULL,
    `width` float NOT NULL,
    `weight` float NOT NULL,
    `expiration_rate` float NOT NULL,
    `freezing_rate` float NOT NULL,
    `recom_freez_temp` float NOT NULL,
    `seller_id` int NULL,
    `product_type_id` int NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_products_product_code` (`product_code`),
    CONSTRAINT `fk_products_seller_id` FOREIGN KEY (`seller_id`) REFERENCES `sellers` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = UTF8MB4;

-- table `employees`
CREATE TABLE `employees` (
    `id` int NOT NULL AUTO_INCREMENT,
    `card_number_id` int NOT NULL,
    `first_name` varchar(50) NOT NULL,
    `last_name` varchar(50) NOT NULL,
    `warehouse_id` int NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_employees_card_number_id` (`card_number_id`),
    CONSTRAINT `fk_employees_warehouse_id` FOREIGN KEY (`warehouse_id`) REFERENCES `warehouses` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = UTF8MB4;

-- table `buyers`
CREATE TABLE `buyers` (
    `id` int NOT NULL AUTO_INCREMENT,
    `card_number_id` int NOT NULL,
    `first_name` varchar(50) NOT NULL,
    `last_name` varchar(50) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_buyers_card_number_id` (`card_number_id`)
) ENGINE = InnoDB DEFAULT CHARSET = UTF8MB4;

-- table `product_batches`
CREATE TABLE `product_batches` (
    `id` int NOT NULL AUTO_INCREMENT,
    `batch_number` int NOT NULL,
    `due_date` date NOT NULL,
    `minimum_temperature` float NOT NULL,
    `current_temperature` float NOT NULL,
    `initial_quantity` int NOT NULL,
    `current_quantity` int NOT NULL,
    `manufacturing_date` date NOT NULL,
    `manufacturing_hour` int NOT NULL,
    `section_id` int NOT NULL,
    `product_id` int NOT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_product_batches_section_id` FOREIGN KEY (`section_id`) REFERENCES `sections` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_product_batches_product_id` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = UTF8MB4;

-- table `inbound_orders`
CREATE TABLE `inbound_orders` (
    `id` int NOT NULL AUTO_INCREMENT,
    `order_number` int NOT NULL,
    `order_date` date NOT NULL,
    `warehouse_id` int NOT NULL,
    `employee_id` int NOT NULL,
    `product_batch_id` int NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_inbound_orders_order_number` (`order_number`),
    CONSTRAINT `fk_inbound_orders_warehouse_id` FOREIGN KEY (`warehouse_id`) REFERENCES `warehouses` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_inbound_orders_employee_id` FOREIGN KEY (`employee_id`) REFERENCES `employees` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_inbound_orders_product_batch_id` FOREIGN KEY (`product_batch_id`) REFERENCES `product_batches` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = UTF8MB4;

-- table `product_records`
CREATE TABLE `product_records` (
    `id` int NOT NULL AUTO_INCREMENT,
    `last_update_date` datetime NOT NULL,
    `purchase_price` float NOT NULL,
    `sale_price` float NULL,
    `product_id` int NOT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_product_records_product_id` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = UTF8MB4;

-- table `purchase_orders`
CREATE TABLE `purchase_orders` (
    `id` int NOT NULL AUTO_INCREMENT,
    `order_number` int NOT NULL,
    `order_date` date NOT NULL,
    `tracking_code` varchar(25) NOT NULL,
    `buyer_id` int NOT NULL,
    `product_record_id` int NOT NULL,
    `order_status_id` int NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_purchase_orders_order_number` (`order_number`),
    CONSTRAINT `fk_purchase_orders_buyer_id` FOREIGN KEY (`buyer_id`) REFERENCES `buyers` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_purchase_orders_product_record_id` FOREIGN KEY (`product_record_id`) REFERENCES `product_records` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = UTF8MB4;