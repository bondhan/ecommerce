CREATE TABLE `products`
(
    `id`          bigint         NOT NULL AUTO_INCREMENT,
    `sku`         varchar(256)   NOT NULL UNIQUE,
    `name`        varchar(256)   NOT NULL,
    `stock`       int            not null default '0',
    `price`       decimal(14, 2) not null default '0',
    `category_id` bigint         NOT NULL,
    `discount_id` bigint                  DEFAULT NULL,
    `created_at`  timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at`  timestamp      NULL,
    PRIMARY KEY (`id`),
    INDEX (`sku`, `deleted_at`, `discount_id`, `category_id`),
    CONSTRAINT products_discount_id_discounts_id FOREIGN KEY (discount_id) REFERENCES discounts (id),
    CONSTRAINT products_category_id_categories_id FOREIGN KEY (category_id) REFERENCES categories (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;