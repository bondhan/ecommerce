CREATE TABLE `products`
(
    `id`          bigint       NOT NULL AUTO_INCREMENT,
    `name`        varchar(256) NOT NULL,
    `stock`       int          not null default '0',
    `price`       int          not null default '0',
    `image`       text null default NULL,
    `category_id` bigint       NOT NULL,
    `discount_id` bigint                DEFAULT NULL,
    `created_at`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at`  timestamp NULL,
    PRIMARY KEY (`id`),
    INDEX ( `name`, `deleted_at`, `discount_id`, `category_id`),
    CONSTRAINT products_discount_id_discounts_id FOREIGN KEY (discount_id) REFERENCES discounts (id),
    CONSTRAINT products_category_id_categories_id FOREIGN KEY (category_id) REFERENCES categories (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;