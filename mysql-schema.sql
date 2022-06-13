create database if not exists go-pulsar-manager default charset utf8mb4 collate utf8mb4_0900_ai_ci;
use go-pulsar-manager;
CREATE TABLE `environment`
(
    id             bigint  NOT NULL AUTO_INCREMENT NOT NULL,
    name           varchar(50) NOT NULL,
    url            varchar(50) NOT NULL,
    description    varchar(50) DEFAULT NULL,
    create_time    datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    create_user    bigint  NULL,
    update_time    datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    update_user    bigint  NULL,
    PRIMARY KEY (`id`)
);