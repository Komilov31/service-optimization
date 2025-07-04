-- +goose Up
CREATE TABLE IF NOT EXISTS deliveries (
    delivery_uid UUID PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    phone TEXT NOT NULL,
    zip VARCHAR(10) NOT NULL,
    city TEXT NOT NULL,
    address TEXT NOT NULL,
    region TEXT NOT NULL,
    email VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS payments (
    payment_uid UUID PRIMARY KEY NOT NULL,
    transaction TEXT NOT NULL,
    request_id TEXT,
    currency VARCHAR(3) NOT NULL,
    provider TEXT NOT NULL,
    amount INT NOT NULL,
    payment_dt BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM now())::BIGINT,
    bank TEXT NOT NULL,
    delivery_cost INT NOT NULL,
    goods_total INT NOT NULL,
    custom_fee INT NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS items (
    item_uid UUID PRIMARY KEY NOT NULL,
    chrt_id BIGINT NOT NULL,
    track_number TEXT NOT NULL,
    price BIGINT NOT NULL,
    rid TEXT NOT NULL,
    name TEXT NOT NULL,
    sale INT NOT NULL,
    size VARCHAR(3) NOT NULL,
    total_price BIGINT NOT NULL,
    nm_id BIGINT NOT NULL,
    brand TEXT NOT NULL,
    status INT NOT NULL
);

CREATE TABLE IF NOT EXISTS orders (
    order_uid UUID PRIMARY KEY NOT NULL,
    track_number TEXT NOT NULL,
    entry TEXT NOT NULL,
    delivery_uid UUID NOT NULL,
    payment_uid UUID NOT NULL,
    locale TEXT NOT NULL,
    internal_signature TEXT,
    customer_id TEXT NOT NULL,
    delivery_service TEXT NOT NULL,
    shardkey TEXT NOT NULL,
    sm_id INT NOT NULL,
    date_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    oof_shard TEXT NOT NULL,

    FOREIGN KEY (delivery_uid) REFERENCES deliveries(delivery_uid),
    FOREIGN KEY (payment_uid) REFERENCES payments(payment_uid)
);

CREATE TABLE IF NOT EXISTS order_items (
    id SERIAL NOT NULL PRIMARY KEY,
    order_uid UUID NOT NULL,
    item_uid UUID NOT NULL,

    FOREIGN KEY (order_uid) REFERENCES orders(order_uid),
    FOREIGN KEY (item_uid) REFERENCES items(item_uid)
);

-- +goose Down
DROP TABLE IF EXISTS orders CASCADE;
DROP TABLE IF EXISTS deliveries CASCADE;
DROP TABLE IF EXISTS payments CASCADE;
DROP TABLE IF EXISTS items CASCADE;
DROP TABLE IF EXISTS order_items CASCADE;
