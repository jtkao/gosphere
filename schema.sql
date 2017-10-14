CREATE DATABASE `sphere`;

USE `sphere`;

CREATE TABLE `food` (  
	`id` INT(4) AUTO_INCREMENT PRIMARY KEY, 
	`name` VARCHAR(255) NOT NULL,     
	`date_created` TIMESTAMP
);

