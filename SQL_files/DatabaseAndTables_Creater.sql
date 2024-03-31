-- This Create the main database and It's tables
-- Date: 2024/03/31
-- Creates the CustomersDB database
CREATE DATABASE IF NOT EXISTS CustomersDB;
-- Uses the CustomersDB database 
USE CustomersDB;
-- Creates the customers table
CREATE TABLE IF NOT EXISTS customers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    dog_name VARCHAR(255) NOT NULL,
    dog_breed VARCHAR(255) NOT NULL,
    dog_age INT(5) NOT NULL
);
-- Creates the appointments table
CREATE TABLE IF NOT EXISTS appointments (
    id INT AUTO_INCREMENT PRIMARY KEY,
    customer_id INT NOT NULL,
    appointment_date DATE NOT NULL,
    appointment_time TIME NOT NULL,
    service VARCHAR(255) NOT NULL
);
ALTER TABLE appointments ADD CONSTRAINT fk_customer_id FOREIGN KEY (customer_id) REFERENCES customers(id);

INSERT INTO customers (name, email, phone, dog_name, dog_breed, dog_age) VALUES ('John Doe', 'johndoe@example.com', '123-456-7890', 'Buddy', 'Labrador Retriever', 3);
INSERT INTO appointments (customer_id, appointment_date, appointment_time, service) VALUES (1, '2002-03-31', '22:00:00', 'Grooming');
INSERT INTO appointments (customer_id, appointment_date, appointment_time, service) VALUES (1, '2222-03-31', '22:00:00', 'Grooming');