CREATE TABLE IF NOT EXISTS products (
	id serial primary key,
	name varchar,
	description varchar,
	price decimal,
	quantity integer
);

INSERT INTO products (
    name,
    description,
    price,
    quantity
)
VALUES
    ('T-shirt', 'A comfortable cotton t-shirt', 19.99, 10),
    ('Headphone', 'An useful Headphone.', 99, 5);

