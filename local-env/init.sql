DROP DATABASE IF EXISTS app;
CREATE DATABASE app;
USE app;
CREATE TABLE user
(
    id           BIGINT PRIMARY KEY AUTO_INCREMENT,
    name         VARCHAR(255),
    email        VARCHAR(255) UNIQUE,
    password     VARCHAR(255) UNIQUE,
    phone_number VARCHAR(255),
    created_on   DATETIME
);

CREATE TABLE address
(
    id            BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id       BIGINT,
    address_line1 TEXT,
    address_line2 TEXT,
    city          VARCHAR(255),
    state         VARCHAR(255),
    zip_code      VARCHAR(255),
    country       VARCHAR(255),
    created_on    DATETIME,
    FOREIGN KEY (user_id) REFERENCES user (id)
);
CREATE TABLE category
(
    id         BIGINT PRIMARY KEY AUTO_INCREMENT,
    name       VARCHAR(255),
    created_on DATETIME
);
CREATE TABLE item
(
    id          BIGINT PRIMARY KEY AUTO_INCREMENT,
    name        VARCHAR(255),
    description TEXT,
    price       DECIMAL(10, 2),
    quantity    INT,
    category_id BIGINT,
    created_on  DATETIME,
    FOREIGN KEY (category_id) REFERENCES category (id)
);

CREATE TABLE cart
(
    id         BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id    BIGINT,
    created_on DATETIME,
    FOREIGN KEY (user_id) REFERENCES user (id)
);

CREATE TABLE cart_item
(
    cart_id  BIGINT,
    item_id  BIGINT,
    quantity INT,
    FOREIGN KEY (cart_id) REFERENCES cart (id),
    FOREIGN KEY (item_id) REFERENCES item (id)
);

CREATE TABLE `order`
(
    id              BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id         BIGINT,
    cart_id         BIGINT,
    total_amount    DECIMAL(10, 2),
    payment_status  VARCHAR(255),
    shipping_status VARCHAR(255),
    created_on      DATETIME,
    FOREIGN KEY (user_id) REFERENCES user (id),
    FOREIGN KEY (cart_id) REFERENCES cart (id)
);
INSERT INTO user (name, email, password, phone_number)
VALUES ('User1', 'user1@example.com', 'password1', '123456789'),
       ('User2', 'user2@example.com', 'password2', '987654321'),
       ('User3', 'user3@example.com', 'password3', '456789123'),
       ('User4', 'user4@example.com', 'password4', '789123456');
-- Addresses for User1
INSERT INTO address (user_id, address_line1, address_line2, city, state, zip_code, country, created_on)
VALUES (1, '123 Main St', 'Apt 101', 'City1', 'State1', '12345', 'Country1', NOW()),
       (1, '456 Oak St', 'Suite 202', 'City1', 'State1', '23456', 'Country1', NOW()),
       (1, '789 Pine St', 'Unit 303', 'City1', 'State1', '34567', 'Country1', NOW());

-- Addresses for User2
INSERT INTO address (user_id, address_line1, address_line2, city, state, zip_code, country, created_on)
VALUES (2, '111 Elm St', 'Apt 11', 'City2', 'State2', '54321', 'Country2', NOW()),
       (2, '222 Cedar St', 'Suite 22', 'City2', 'State2', '65432', 'Country2', NOW()),
       (2, '333 Maple St', 'Unit 33', 'City2', 'State2', '76543', 'Country2', NOW());

-- Addresses for User3
INSERT INTO address (user_id, address_line1, address_line2, city, state, zip_code, country, created_on)
VALUES (3, '555 Birch St', 'Apt 55', 'City3', 'State3', '98765', 'Country3', NOW()),
       (3, '666 Walnut St', 'Suite 66', 'City3', 'State3', '87654', 'Country3', NOW()),
       (3, '777 Sycamore St', 'Unit 77', 'City3', 'State3', '76543', 'Country3', NOW());

-- Addresses for User4
INSERT INTO address (user_id, address_line1, address_line2, city, state, zip_code, country, created_on)
VALUES (4, '999 Pineapple St', 'Apt 99', 'City4', 'State4', '23456', 'Country4', NOW()),
       (4, '888 Banana St', 'Suite 88', 'City4', 'State4', '34567', 'Country4', NOW()),
       (4, '777 Mango St', 'Unit 77', 'City4', 'State4', '45678', 'Country4', NOW());

