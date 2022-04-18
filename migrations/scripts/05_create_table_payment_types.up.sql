CREATE TABLE `payment_types`
(
    `id`         bigint       NOT NULL AUTO_INCREMENT,
    `name`       varchar(256) NOT NULL,
    `type`       enum ('CASH', 'E-WALLET', 'EDC'),
    `logo`       TEXT         null     default null,
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` timestamp    NULL,
    PRIMARY KEY (`id`),
    INDEX (`deleted_at`, `type`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;