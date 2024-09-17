CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `firstname` varchar(100) NOT NULL,
  `lastname` varchar(100) NOT NULL,
  `age` int NULL
);

INSERT INTO `user` (`firstname`, `lastname`, `age`)
VALUES ('oliver', 'reese', '52');