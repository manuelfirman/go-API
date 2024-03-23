USE `go_api_db`;

-- TRUNCATE TABLE `sellers`;
-- TRUNCATE TABLE `warehouses`;
-- TRUNCATE TABLE `sections`;
-- TRUNCATE TABLE `products`;
-- TRUNCATE TABLE `employees`;
-- TRUNCATE TABLE `buyers`;
-- TRUNCATE TABLE `product_batches`;

-- DML
INSERT INTO `localities` (`id`, `locality_name`, `province_name`, `country_name`) VALUES
(100, 'City A', 'Province A', 'Country A'),
(102, 'City B', 'Province B', 'Country A'),
(103, 'City C', 'Province B', 'Country A'),
(104, 'City A', 'Province C', 'Country B'),
(105, 'City A', 'Province C', 'Country B'),
(106, 'City C', 'Province C', 'Country C'),
(107, 'City B', 'Province A', 'Country C'),
(108, 'City B', 'Province A', 'Country C'),
(109, 'City A', 'Province A', 'Country B'),
(110, 'City A', 'Province B', 'Country A');

INSERT INTO `carries` (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES
(1, 'Company A', '123 Main St', '123-456-7890', 100),
(2, 'Company B', '456 Elm St', '123-456-7891', 104),
(3, 'Company B', '789 Oak St', '123-456-7892', 110),
(4, 'Company A', '101 Pine St', '123-456-7893', 109),
(5, 'Company A', '102 Maple St', '123-456-7894', 104),
(6, 'Company C', '103 Cedar St', '123-456-7895', 102),
(7, 'Company C', '104 Birch St', '123-456-7896', 104),
(8, 'Company B', '105 Willow St', '123-456-7897', 107),
(9, 'Company A', '106 Cherry St', '123-456-7898', 108),
(10, 'Company C', '107 Walnut St', '123-456-7899', 109);

INSERT INTO `sellers` (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES
(1, 'Company A', '123 Main St', '123-456-7890', 100),
(2, 'Company B', '456 Elm St', '123-456-7891', 103),
(3, 'Company C', '789 Oak St', '123-456-7892', 106),
(4, 'Company D', '101 Pine St', '123-456-7893', 100),
(5, 'Company E', '102 Maple St', '123-456-7894', 102),
(6, 'Company F', '103 Cedar St', '123-456-7895', 102),
(7, 'Company G', '104 Birch St', '123-456-7896', 100),
(8, 'Company H', '105 Willow St', '123-456-7897', 100),
(9, 'Company I', '106 Cherry St', '123-456-7898', 102),
(10, 'Company J', '107 Walnut St', '123-456-7899', 110);

INSERT INTO `warehouses` (`warehouse_code`, `address`, `telephone`, `minimum_capacity`, `minimum_temperature`, `locality_id`) VALUES
('WH01', '200 Warehouse Rd', '234-567-8901', 100, 0, 100),
('WH02', '201 Warehouse Ln', '234-567-8902', 150, -5, 104),
('WH03', '202 Storage Blvd', '234-567-8903', 120, 2, 104),
('WH04', '203 Distribution Ave', '234-567-8904', 200, -2, 103),
('WH05', '204 Inventory St', '234-567-8905', 180, 0, 105),
('WH06', '205 Logistics Way', '234-567-8906', 160, -3, 100),
('WH07', '206 Depot Dr', '234-567-8907', 140, 1, 102),
('WH08', '207 Supply Ct', '234-567-8908', 170, -4, 108),
('WH09', '208 Goods Rd', '234-567-8909', 130, 3, 110),
('WH10', '209 Freight St', '234-567-8910', 190, -1, 107);

INSERT INTO `sections` (`section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id`) VALUES
(1, 0, -5, 50, 20, 100, 1, 1),
(2, -2, -6, 60, 30, 110, 2, 2),
(3, 1, -4, 70, 40, 120, 3, 3),
(4, -3, -7, 80, 50, 130, 4, 4),
(5, 2, -5, 90, 60, 140, 5, 5),
(6, -4, -8, 100, 70, 150, 6, 6),
(7, 3, -6, 110, 80, 160, 7, 7),
(8, -5, -9, 120, 90, 170, 8, 8),
(9, 4, -7, 130, 100, 180, 9, 9),
(10, -6, -10, 140, 110, 190, 10, 10);

INSERT INTO `products` (`product_code`, `description`, `height`, `length`, `width`, `weight`, `expiration_rate`, `freezing_rate`, `recom_freez_temp`, `seller_id`, `product_type_id`) VALUES
('P1001', 'Product 1', 10, 5, 8, 2, 0.1, 0.2, -5, 1, 1),
('P1002', 'Product 2', 12, 6, 9, 2.5, 0.15, 0.25, -6, 2, 2),
('P1003', 'Product 3', 14, 7, 10, 3, 0.2, 0.3, -7, 3, 3),
('P1004', 'Product 4', 16, 8, 11, 3.5, 0.25, 0.35, -8, 4, 4),
('P1005', 'Product 5', 18, 9, 12, 4, 0.3, 0.4, -9, 5, 5),
('P1006', 'Product 6', 20, 10, 13, 4.5, 0.35, 0.45, -10, 6, 6),
('P1007', 'Product 7', 22, 11, 14, 5, 0.4, 0.5, -11, 7, 7),
('P1008', 'Product 8', 24, 12, 15, 5.5, 0.45, 0.55, -12, 8, 8),
('P1009', 'Product 9', 26, 13, 16, 6, 0.5, 0.6, -13, 9, 9),
('P1010', 'Product 10', 28, 14, 17, 6.5, 0.55, 0.65, -14, 10, 10);

INSERT INTO `employees` (`card_number_id`, `first_name`, `last_name`, `warehouse_id`) VALUES
(1001, 'John', 'Doe', 1),
(1002, 'Jane', 'Smith', 2),
(1003, 'Michael', 'Johnson', 3),
(1004, 'Emily', 'Davis', 4),
(1005, 'David', 'Miller', 5),
(1006, 'Sarah', 'Wilson', 6),
(1007, 'Robert', 'Moore', 7),
(1008, 'Jennifer', 'Taylor', 8),
(1009, 'William', 'Anderson', 9),
(1010, 'Jessica', 'Thomas', 10);

INSERT INTO `buyers` (`card_number_id`, `first_name`, `last_name`) VALUES
(1001, 'Alice', 'Brown'),
(1002, 'Mark', 'Jones'),
(1003, 'Linda', 'Garcia'),
(1004, 'Brian', 'Williams'),
(1005, 'Susan', 'Martinez'),
(1006, 'Richard', 'Lee'),
(1007, 'Karen', 'Harris'),
(1008, 'Steven', 'Clark'),
(1009, 'Betty', 'Lopez'),
(1010, 'Edward', 'Gonzalez');

INSERT INTO `product_batches` (`batch_number`, `due_date`, `minimum_temperature`, `current_temperature`, `initial_quantity`, `current_quantity`, `manufacturing_date`, `manufacturing_hour`, `section_id`, `product_id`) VALUES
(1, '2021-12-31', -5, 0, 200, 100, '2021-01-01', 1, 1, 3),
(2, '2021-12-31', -4, 5, 100, 50, '2021-01-01', 12, 1, 1),
(3, '2021-12-31', -2, 5, 100, 33, '2021-01-01', 16, 5, 2),
(4, '2021-12-31', 0, 12, 56, 21, '2021-01-01', 13, 6, 6),
(5, '2021-12-31', 5, 22, 67, 120, '2021-01-01', 6, 8, 6),
(6, '2021-12-31', -10, 0, 12, 20, '2021-01-01', 6, 8, 8),
(7, '2021-12-31', 2, 12, 111, 222, '2021-01-01', 7, 8, 9),
(8, '2021-12-31', 5, 6, 23, 50, '2021-01-01', 5, 9, 10),
(9, '2021-12-31', -2, 4, 100, 100, '2021-01-01', 8, 3, 4),
(10, '2021-12-31', -1, 5, 150, 200, '2021-01-01', 5, 10, 10);

INSERT INTO `inbound_orders` (`order_number`, `order_date`, `warehouse_id`, `employee_id`, `product_batch_id`) VALUE
(1, '2021-04-12', 1, 1, 1),
(2, '2021-06-01', 2, 3, 5),
(3, '2021-01-12', 6, 9, 7),
(4, '2022-12-12', 8, 3, 7),
(5, '2021-06-13', 9, 9, 9),
(6, '2021-07-23', 1, 1, 1),
(7, '2021-03-23', 4, 2, 1),
(8, '2022-12-12', 7, 2, 1),
(9, '2021-05-05', 6, 3, 9),
(10, '2021-04-12', 7, 2, 10);

INSERT INTO `product_records` (`last_update_date`, `purchase_price`, `sale_price`, `product_id`) VALUES
('2021-01-01', 10, 5, 1),
('2021-01-01', 15, 10, 2),
('2021-01-01', 16, NULL, 3),
('2021-01-01', 19, 15, 4),
('2021-01-01', 15, 12, 5),
('2021-01-01', 30, 25, 6),
('2021-01-01', 50, NULL, 7),
('2021-01-01', 12, 6, 8),
('2021-01-01', 6, 1, 9),
('2021-01-01', 10, 6, 10);

INSERT INTO `purchase_orders` (`order_number`, `order_date`, `tracking_code`, `buyer_id`, `product_record_id`, `order_status_id`) VALUES
(1, '2021-01-01', 'ABC123', 1, 1, 1),
(2, '2021-01-01', 'ABC124', 2, 2, 2),
(3, '2021-01-01', 'ABC123', 2, 2, 1),
(4, '2021-01-01', 'ABC123', 3, 3, 3),
(5, '2021-01-01', 'ABC125', 8, 6, 2),
(6, '2021-01-01', 'ABC125', 9, 6, 5),
(7, '2021-01-01', 'ABC129', 10, 7, 3),
(8, '2021-01-01', 'ABC130', 1, 1, 2),
(9, '2021-01-01', 'ABC131', 2, 2, 1),
(10, '2021-01-01', 'ABC132', 3, 3, 4);