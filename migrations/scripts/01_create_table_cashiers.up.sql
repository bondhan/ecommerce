CREATE TABLE `cashiers`
(
    `id`           bigint       NOT NULL AUTO_INCREMENT,
    `name`         varchar(256) NOT NULL UNIQUE,
    `passcode`     varchar(256) NOT NULL,
    `login_status` enum ('logged_in', 'logged_out', 'blocked'),
    `created_at`   timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`   timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at`   timestamp    NULL,
    PRIMARY KEY (`id`),
    INDEX (`login_status`, `deleted_at`, `name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;