CREATE TABLE `environment`
(
    id             bigint  NOT NULL AUTO_INCREMENT NOT NULL,
    name           varchar(50) NOT NULL,
    url            varchar(50) NOT NULL,
    description    varchar(50) DEFAULT NULL,
    PRIMARY KEY (`id`)
);