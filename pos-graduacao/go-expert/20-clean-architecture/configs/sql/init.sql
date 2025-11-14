CREATE TABLE orders (
    id VARCHAR(255) NOT NULL,
    price FLOAT NOT NULL,
    tax FLOAT NOT NULL,
    final_price FLOAT NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO orders (id, price, tax, final_price) VALUES 
('order-ex-1', 100.0, 10.0, 110.0),
('order-ex-2', 250.5, 25.05, 275.55),
('order-ex-3', 75.0, 7.5, 82.5),
('order-ex-4', 500.0, 50.0, 550.0),
('order-ex-5', 199.99, 19.99, 219.98);

SELECT 'Orders table created and populated with sample data' as message;
SELECT COUNT(*) as total_orders FROM orders;
