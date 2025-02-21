-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user" (
                        id SERIAL PRIMARY KEY,
                        fullname TEXT,
                        login TEXT NOT NULL,
                        password TEXT NOT NULL,
                        is_admin BOOLEAN DEFAULT FALSE
);

CREATE TABLE goods (
                       id SERIAL PRIMARY KEY,
                       name TEXT NOT NULL,
                       description TEXT,
                       price INTEGER NOT NULL,
                       count INTEGER NOT NULL DEFAULT 0,
                       created_by_user_id INTEGER REFERENCES "user"(id) ON DELETE SET NULL,
                       created_at TIMESTAMP DEFAULT NOW(),
                       updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE stock (
                       id SERIAL PRIMARY KEY,
                       address TEXT NOT NULL
);

CREATE TABLE goods_stock (
                             goods_id INTEGER REFERENCES goods(id) ON DELETE CASCADE,
                             stock_id INTEGER REFERENCES stock(id) ON DELETE CASCADE,
                             goods_count INTEGER DEFAULT 0,
                             PRIMARY KEY (goods_id, stock_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE goods_stock;
DROP TABLE stock;
DROP TABLE goods;
DROP TABLE "user";

-- +goose StatementEnd
