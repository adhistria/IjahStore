-- migrate:up

CREATE TABLE IncomingProducts (
    id INTEGER PRIMARY KEY,
    total_order INTEGER,
    total_received_order INTEGER,
    purchase_price INTEGER,
    total_purchase_price INTEGER,
    receipt_number VARCHAR(100),
    notes VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    product_id VARCHAR(255) NOT NULL,
    FOREIGN KEY(product_id) REFERENCES Products(SKU)
);

-- migrate:down

DROP TABLE IncomingProducts;
