CREATE DATABASE `sphere`;

USE `sphere`;

CREATE TABLE `wishlist` (  
	`id` INT(4) AUTO_INCREMENT PRIMARY KEY, 
	`name` VARCHAR(255) NOT NULL,     
	`date_created` TIMESTAMP
);

INSERT INTO `wishlist` (`name`) VALUES 
("Westvale Abbey"),
("Torment of Hailfire"),
("Gifted Aetherborn"),
("Geth, Lord of the Vault"),
("Grave Pact"),
("Demon of Death's Gate");
