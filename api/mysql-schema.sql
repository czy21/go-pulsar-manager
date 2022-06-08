CREATE TABLE `cluster`
(
    id             bigint  NOT NULL AUTO_INCREMENT NOT NULL,
    name           varchar(50) NOT NULL,
    url            varchar(50) NOT NULL,
    description    varchar(50) DEFAULT NULL,
    create_time    datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    create_user    varchar(36) NULL,
    update_time    datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    update_user    varchar(36) NULL,
    deleted        bit(1)      NOT NULL DEFAULT b'0',
    PRIMARY KEY (`id`)
);