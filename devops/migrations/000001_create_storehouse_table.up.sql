CREATE TABLE IF NOT EXISTS storehouse(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    is_available BOOLEAN NOT NULL
);
INSERT INTO storehouse (name, is_available)
VALUES
    ('Stock 1', true),
    ('Stock 2', false),
    ('Stock 3', false);