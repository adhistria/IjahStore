CREATE TABLE schema_migrations (version varchar(255) primary key);
CREATE TABLE Products (
    SKU VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255),
    total INTEGER
);
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
-- Dbmate schema migrations
INSERT INTO schema_migrations (version) VALUES
  ('20190901125905'),
  ('20190901142839'),
  ('20190901160942');
