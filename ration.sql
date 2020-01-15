-- Adminer 4.7.0 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DROP TABLE IF EXISTS `ration`;
CREATE TABLE `ration` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `packet_type` varchar(11) NOT NULL,
  `packet_content` varchar(100) DEFAULT NULL,
  `calories` int(5) NOT NULL DEFAULT '0',
  `expiry_date` date DEFAULT NULL,
  `quantity` int(5) NOT NULL,
  `created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `ration` (`id`, `packet_type`, `packet_content`, `calories`, `expiry_date`, `quantity`, `created_at`, `updated_at`) VALUES
(76,	'food',	'MRE',	1000,	'2020-01-16',	0,	'2020-01-15 08:20:59',	'2020-01-15 08:20:59'),
(77,	'food',	'Chips',	500,	'2020-01-24',	0,	'2020-01-15 08:21:12',	'2020-01-15 08:21:12'),
(78,	'food',	'Biscuits',	1000,	'2020-01-23',	0,	'2020-01-15 08:21:20',	'2020-01-15 08:21:20'),
(79,	'water',	'',	0,	NULL,	3,	'2020-01-15 08:21:24',	'2020-01-15 08:21:24');

-- 2020-01-15 09:14:49
