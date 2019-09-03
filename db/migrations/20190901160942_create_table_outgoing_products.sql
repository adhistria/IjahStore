-- migrate:up

CREATE TABLE OutgoingProducts (
    id INTEGER PRIMARY KEY,
    sold_amount INTEGER,
    selling_price INTEGER,
    total_selling_price INTEGER,
    notes VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    product_id VARCHAR(255) NOT NULL,
    FOREIGN KEY(product_id) REFERENCES Products(SKU)
);

-- migrate:down

DROP TABLE OutgoingProducts;
