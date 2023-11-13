CREATE TABLE IF NOT EXISTS product(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    size VARCHAR(10),
    quantity INT DEFAULT 0,
    reserved INT DEFAULT 0,
    storehouse_id INT,
    FOREIGN KEY (storehouse_id) REFERENCES storehouse(id)
);
ALTER TABLE product
    ADD CONSTRAINT positive_quantity_constraint CHECK (quantity >= 0);
ALTER TABLE product
    ADD CONSTRAINT positive_reserved_constraint CHECK (reserved >= 0);
INSERT INTO product (name, size, quantity, reserved, storehouse_id)
VALUES
    ('Product 1', 'Small', 100, 0,1),
    ('Product 2', 'Medium', 150, 0,1),
    ('Product 3', 'Large', 200, 0,2),
    ('Product 4', 'Small', 75, 0,2),
    ('Product 5', 'Medium', 120, 4, 1),
    ('Product 6', 'Large', 180, 20,2),
    ('Product 7', 'Small', 90, 5, 2),
    ('Product 8', 'Medium', 110, 0, 1),
    ('Product 9', 'Large', 160, 0, 1),
    ('Product 10', 'Small', 50, 0, 1),
    ('Product 11', 'Medium', 70, 0, 2),
    ('Product 12', 'Large', 220, 15, 1);