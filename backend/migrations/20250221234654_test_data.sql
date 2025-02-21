-- +goose Up
-- +goose StatementBegin
INSERT INTO "user" (id, fullname, login, password, is_admin) VALUES
                                                                 (2, 'aaa', 'bbb', 'aaa', TRUE),
                                                                 (3, 'a', 'a', 'a', TRUE),
                                                                 (4, 'a', 'a', 'c', FALSE),
                                                                 (6, 'semyon t', 'login', 'password', FALSE),
                                                                 (7, 'ss', 'ss', 'ss', FALSE),
                                                                 (8, 'ss', 'ss', 'ss', FALSE),
                                                                 (9, 'ss', 'ss', 'ss', FALSE),
                                                                 (10, 'xx', 'cc', 'vv', TRUE),
                                                                 (11, 'test', 'test', 'test', TRUE);

INSERT INTO goods (id, name, description, price, count, created_by_user_id, created_at, updated_at) VALUES
                                                                                                        (4, 'name', 'desc', 100, 1, 2, '2024-06-17 01:21:37', '2024-06-17 01:21:37'),
                                                                                                        (5, 'name', 'desc', 100, 2, 2, '2024-06-17 01:23:22', '2024-06-17 01:23:22'),
                                                                                                        (6, 'aaa', 'bbb', 123, 2, 2, '2024-06-17 01:28:54', '2024-06-17 01:28:54'),
                                                                                                        (9, 'aaa', 'bbb', 222222, 2, 3, '2024-06-17 02:07:15', '2024-06-17 02:07:15'),
                                                                                                        (11, 'aaa', 'bbb', 23, 2, 11, '2024-06-17 09:53:41', '2024-06-17 09:53:41'),
                                                                                                        (13, 'aaa', 'bbb', 23, 2, 3, '2024-06-17 09:57:41', '2024-06-17 09:57:41'),
                                                                                                        (14, 'aaa', 'bbb', 23, 2, 3, '2024-06-17 09:57:45', '2024-06-17 09:57:45'),
                                                                                                        (15, 'test', 'test', 213, 2, 2, '2024-06-17 10:54:39', '2024-06-17 10:54:39'),
                                                                                                        (19, 'dd', 'dsds', 123, 5, 2, '2024-06-17 15:01:12', '2024-06-17 15:01:12'),
                                                                                                        (21, 'test file', 'test file', 22, 1, 11, '2024-06-17 15:05:54', '2024-06-17 15:05:54');

INSERT INTO stock (id, address) VALUES
    (3, 'nnn');

INSERT INTO goods_stock (goods_id, stock_id, goods_count) VALUES
    (14, 3, 100);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
