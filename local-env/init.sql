DROP DATABASE IF EXISTS app;
CREATE DATABASE app;
USE app;
CREATE TABLE user
(
    id           BIGINT PRIMARY KEY AUTO_INCREMENT,
    name         VARCHAR(255),
    email        VARCHAR(255) UNIQUE, -- Assuming email for login
    password     VARCHAR(255),
    phone_number VARCHAR(255),
    created_on   DATETIME,
    last_login   DATETIME
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

CREATE TABLE order
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