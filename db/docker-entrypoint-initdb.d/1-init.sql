CREATE DATABASE IF NOT EXISTS `counterdb`;
USE counterdb;
CREATE TABLE IF NOT EXISTS `counter`
(
    `id`        INT unsigned NOT NULL AUTO_INCREMENT,
    `value`     INT unsigned NOT NULL DEFAULT '0' COMMENT 'current counter value',
    `createdAt` TIMESTAMP,
    `updatedAt` TIMESTAMP ON UPDATE CURRENT_TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (`id`)
);

CREATE DATABASE IF NOT EXISTS `counterdb_test`;
USE counterdb_test;
CREATE TABLE IF NOT EXISTS `counter`
(
    `id`        INT unsigned NOT NULL AUTO_INCREMENT,
    `value`     INT unsigned NOT NULL DEFAULT '0' COMMENT 'current counter value',
    `createdAt` TIMESTAMP,
    `updatedAt` TIMESTAMP ON UPDATE CURRENT_TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (`id`)
);