CREATE TABLE `orders`
(
    `id`              bigint         NOT NULL AUTO_INCREMENT,
    `totalPrice`      decimal(14, 2) not null default '0',
    `totalPaid`       decimal(14, 2) not null default '0',
    `totalReturn`     decimal(14, 2) not null default '0',
    `cashier_id`      bigint         NOT NULL,
    `payment_type_id` bigint         NOT NULL,
    `created_at`      timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`      timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at`      timestamp      NULL,
    PRIMARY KEY (`id`),
    INDEX (`deleted_at`, `cashier_id`, `payment_type_id`),
    CONSTRAINT orders_cashier_id_cashiers_id FOREIGN KEY (cashier_id) REFERENCES cashiers (id),
    CONSTRAINT orders_payment_type_id_payment_types_id FOREIGN KEY (payment_type_id) REFERENCES payment_types (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;