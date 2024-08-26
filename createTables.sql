CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL
);

CREATE TABLE Products (
    id SERIAL PRIMARY KEY,
    serial TEXT NOT NULL,
    brand VARCHAR(255) NOT NULL,
    model VARCHAR(255) NOT NULL,
    institution_id INT,
    responsible TEXT,
    owner TEXT,
    status TEXT,
    FOREIGN KEY (institution_id) REFERENCES institutions(id)
);

CREATE TABLE Cycles (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    year INT NOT NULL,
    month INT NOT NULL,
    cycle INT,
    FOREIGN KEY (product_id) REFERENCES Products(id)
);

CREATE TABLE User_Product_Assignments (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    product_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users(id),
    FOREIGN KEY (product_id) REFERENCES Products(id)
);

CREATE TABLE institutions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL
);

