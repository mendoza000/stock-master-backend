-- Active: 1668828896969@@127.0.0.1@3306@stockmaster
CREATE DATABASE stockmaster;
USE stockmaster;
CREATE TABLE products (
    id INT UNIQUE AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(100) UNIQUE NOT NULL,
    brand VARCHAR(100) NOT NULL,
    stock BOOLEAN NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    details VARCHAR(150) NOT NULL
);
CREATE TABLE ranges (
    id INT UNIQUE AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL
);
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    mail VARCHAR(200) UNIQUE NOT NULL,
    password VARCHAR(255) UNIQUE NOT NULL,
    user_range VARCHAR(15) NOT NULL
);
    -- FOREIGN KEY (user_range) REFERENCES ranges(id)

ALTER TABLE products
ADD amount INT NOT NULL;

INSERT INTO products (title, brand, stock, price, details) VALUES
    ('Teclado inalámbrico', 'Logitech', 10, 19.99, 'Teclado inalámbrico con retroiluminación y conectividad Bluetooth'),
    ('Mouse óptico', 'Microsoft', 1, 9.99, 'Mouse óptico con cable y resolución de 1200 dpi'),
    ('Monitor LED', 'Samsung', 1, 149.99, 'Monitor LED de 24 pulgadas con resolución Full HD');

SELECT * FROM products;
DELETE FROM products WHERE id = 0;

INSERT INTO ranges (name) VALUES ("ADMIN"), ("USER"), ("SELLER");
SELECT * FROM ranges;
