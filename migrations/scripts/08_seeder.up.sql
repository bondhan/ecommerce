INSERT INTO ecommerce.cashiers (id, name, passcode, login_status, created_at, updated_at, deleted_at) VALUES(1, 'Kasir 2', '123456', 'logged_out', '2022-04-19 17:44:03', '2022-04-19 17:44:03', NULL);

INSERT INTO ecommerce.categories (id, name, created_at, updated_at, deleted_at) VALUES(1, 'Electronics', '2022-04-19 16:37:36', '2022-04-19 16:37:36', NULL);
INSERT INTO ecommerce.categories (id, name, created_at, updated_at, deleted_at) VALUES(2, 'Snack', '2022-04-19 16:37:36', '2022-04-19 16:37:36', NULL);
INSERT INTO ecommerce.categories (id, name, created_at, updated_at, deleted_at) VALUES(3, 'Clothes', '2022-04-19 16:37:36', '2022-04-19 16:37:36', NULL);

INSERT INTO ecommerce.payment_types (id, name, `type`, logo, created_at, updated_at, deleted_at) VALUES(1, 'Cash', 'CASH', NULL, '2022-04-19 17:44:17', '2022-04-19 17:44:17', NULL);

INSERT INTO ecommerce.discounts (id, `type`, qty, `result`, expired_at, created_at, updated_at, deleted_at) VALUES(1, 'PERCENT', 1, 15, '2022-04-30 00:00:00', '2022-04-19 16:38:35', '2022-04-19 16:38:35', NULL);
INSERT INTO ecommerce.discounts (id, `type`, qty, `result`, expired_at, created_at, updated_at, deleted_at) VALUES(2, 'BUY_N', 2, 8000, '2022-04-30 00:00:00', '2022-04-19 16:39:14', '2022-04-19 16:39:14', NULL);
INSERT INTO ecommerce.discounts (id, `type`, qty, `result`, expired_at, created_at, updated_at, deleted_at) VALUES(3, 'BUY_N', 1, 2000, '2022-04-30 00:00:00', '2022-04-19 16:59:27', '2022-04-19 16:59:27', NULL);

INSERT INTO ecommerce.products (id, name, stock, price, image, category_id, discount_id, created_at, updated_at, deleted_at) VALUES(1, 'Chiki ball', 100, 7000, 'https://images.tokopedia.net/img/cache/500-square/product-1/2020/2/13/35604504/35604504_105cab00-a047-41a9-beaa-a27e755b61f2_1100_1100', 2, 1, '2022-04-19 16:38:35', '2022-04-19 16:38:35', NULL);
INSERT INTO ecommerce.products (id, name, stock, price, image, category_id, discount_id, created_at, updated_at, deleted_at) VALUES(2, 'Monitor', 100, 9000, 'https://images.tokopedia.net/img/cache/500-square/product-1/2020/2/13/35604504/35604504_105cab00-a047-41a9-beaa-a27e755b61f2_1100_1100', 1, 2, '2022-04-19 16:39:14', '2022-04-19 16:39:14', NULL);
INSERT INTO ecommerce.products (id, name, stock, price, image, category_id, discount_id, created_at, updated_at, deleted_at) VALUES(3, 'Baju', 110, 4000, 'https://images.tokopedia.net/img/cache/500-square/product-1/2020/2/13/35604504/35604504_105cab00-a047-41a9-beaa-a27e755b61f2_1100_1100', 3, NULL, '2022-04-19 16:42:38', '2022-04-19 16:42:38', NULL);
INSERT INTO ecommerce.products (id, name, stock, price, image, category_id, discount_id, created_at, updated_at, deleted_at) VALUES(6, 'Laptop', 100, 3000, 'https://images.tokopedia.net/img/cache/500-square/product-1/2020/2/13/35604504/35604504_105cab00-a047-41a9-beaa-a27e755b61f2_1100_1100', 1, 3, '2022-04-19 16:59:27', '2022-04-19 16:59:27', NULL);

