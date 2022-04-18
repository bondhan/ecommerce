CREATE TABLE `discounts`
(
    `id`         bigint         NOT NULL AUTO_INCREMENT,
    `type`       enum ('BUY_N', 'PERCENT'),
    `result`     decimal(14, 2) not null default '0',
    `expired_at` timestamp      NOT NULL,
    `created_at` timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` timestamp      NULL,
    PRIMARY KEY (`id`),
    INDEX (`type`, `deleted_at`, `expired_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;