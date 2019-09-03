-- migrate:up
CREATE TABLE Products (
    SKU VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255),
    total INTEGER
);

-- migrate:down

DROP TABLE Products;
