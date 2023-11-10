CREATE TABLE IF NOT EXISTS product(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    size VARCHAR(10),
    code VARCHAR(20) NOT NULL,
    quantity INT DEFAULT 0,
    reserved INT DEFAULT 0,
    stock_id INT
);
ALTER TABLE product
    ADD CONSTRAINT positive_quantity_constraint CHECK (quantity >= 0);
ALTER TABLE product
    ADD CONSTRAINT positive_reserved_constraint CHECK (reserved >= 0);
ALTER TABLE product
    ADD CONSTRAINT unique_code_per_stock UNIQUE (stock_id, code);
INSERT INTO product (name, size, code, quantity, stock_id)
VALUES
    ('Product 1', 'Small', 'P0001', 100, 1),
    ('Product 2', 'Medium', 'P0002', 150, 1),
    ('Product 3', 'Large', 'P0003', 200, 2),
    ('Product 4', 'Small', 'P0004', 75, 2),
    ('Product 5', 'Medium', 'P0005', 120, 1),
    ('Product 6', 'Large', 'P0006', 180, 2),
    ('Product 7', 'Small', 'P0007', 90, 2),
    ('Product 8', 'Medium', 'P0008', 110, 1),
    ('Product 9', 'Large', 'P0009', 160, 1),
    ('Product 10', 'Small', 'P0010', 50, 1),
    ('Product 11', 'Medium', 'P0011', 70, 2),
    ('Product 12', 'Large', 'P0012', 220, 1);