# linh-phan-golang-fresher-camp
CREATE TABLE `Users` (
  `User_id` int PRIMARY KEY AUTO_INCREMENT,
  `User_name` varchar(255),
  `User_DOB` date,
  `User_email` varchar(255),
  `User_phone` varchar(255),
  `User_date_created_at` date
);

CREATE TABLE `Restaurants` (
  `Restaurant_id` int PRIMARY KEY AUTO_INCREMENT,
  `Restaurants_name` varchar(255),
  `Restaurant_email` varchar(255),
  `Restaurant_phone` varchar(255),
  `Restaurant_address` varchar(255),
  `Restaurant_open` datetime,
  `Restaurant_close` datetime,
  `Restaurants_created_at` date
);

CREATE TABLE `Products` (
  `Product_id` int PRIMARY KEY AUTO_INCREMENT,
  `Product_name` varchar(255),
  `Product_type` int,
  `Product_price` int,
  `Product_of_restaurant` int,
  `Product_combo` int,
  `Product_deal` int,
  `Product_sale` int,
  `Product_status` boolean
);

CREATE TABLE `TypeOfProducts` (
  `TypeOfProduct_id` int PRIMARY KEY AUTO_INCREMENT,
  `TypeOfProduct_name` varchar(255)
);

CREATE TABLE `Combos` (
  `Combo_id` int PRIMARY KEY AUTO_INCREMENT,
  `Combo_name` varchar(255),
  `Combo_price` int
);

CREATE TABLE `Deals` (
  `Deal_id` int PRIMARY KEY AUTO_INCREMENT,
  `Deal_name` varchar(255),
  `Deal_price` int
);

CREATE TABLE `Sales` (
  `Sale_id` int PRIMARY KEY AUTO_INCREMENT,
  `Sale_name` varchar(255),
  `Sale_percent` int
);

ALTER TABLE `TypeOfProducts` ADD FOREIGN KEY (`TypeOfProduct_id`) REFERENCES `Products` (`Product_type`);

ALTER TABLE `Products` ADD FOREIGN KEY (`Product_of_restaurant`) REFERENCES `Restaurants` (`Restaurant_id`);

ALTER TABLE `Products` ADD FOREIGN KEY (`Product_sale`) REFERENCES `Sales` (`Sale_id`);

ALTER TABLE `Products` ADD FOREIGN KEY (`Product_deal`) REFERENCES `Deals` (`Deal_id`);

ALTER TABLE `Products` ADD FOREIGN KEY (`Product_combo`) REFERENCES `Combos` (`Combo_id`);
