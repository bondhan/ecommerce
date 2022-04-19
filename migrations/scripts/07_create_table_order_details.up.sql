CREATE TABLE `order_details`
(
    `id`               bigint       NOT NULL AUTO_INCREMENT,
    `name`             varchar(256) NOT NULL,
    `price`            int          not null default '0',
    `qty`              int          not null default '0',
    `order_id`         bigint       NOT NULL,
    `product_id`       bigint       NOT NULL,
    `discount_id`      bigint       NOT NULL,
    `totalNormalPrice` int          not null default '0',
    `totalFinalPrice`  int          not null default '0',
    `created_at`       timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`       timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at`       timestamp NULL,
    PRIMARY KEY (`id`),
    INDEX (`deleted_at`, `product_id`, `order_id`),
    CONSTRAINT order_details_order_id_orders_id FOREIGN KEY (order_id) REFERENCES orders (id),
    CONSTRAINT order_details_product_id_products_id FOREIGN KEY (product_id) REFERENCES products (id),
    CONSTRAINT order_details_discount_id_discounts_id FOREIGN KEY (discount_id) REFERENCES discounts (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;