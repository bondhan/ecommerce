CREATE TABLE `categories`
(
    `id`         bigint       NOT NULL AUTO_INCREMENT,
    `name`       varchar(256) NOT NULL UNIQUE,
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` timestamp    NULL,
    PRIMARY KEY (`id`),
    INDEX (`name`, `deleted_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;