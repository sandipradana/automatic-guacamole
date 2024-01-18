-- Create Users table
CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

-- Insert Users
INSERT INTO Users ("name",email,"password") VALUES('User','user@localhost','$2a$14$7mxKEL.BpXDFqcy.o9McMuTtTAxHEMim0O6EbLmf38F5RR8L1ZO6q');


-- Create Products table
CREATE TABLE Products (
    ID    SERIAL PRIMARY KEY,
    Name  VARCHAR(255),
    Price DOUBLE PRECISION
);

-- Create ProductAddonGroups table
CREATE TABLE ProductAddonGroups (
    ID        SERIAL PRIMARY KEY,
    ProductsID BIGINT REFERENCES Products(ID),
    Name      VARCHAR(255),
    MaxQuantity INT
);

-- Create ProductAddons table
CREATE TABLE ProductAddons (
    ID             SERIAL PRIMARY KEY,
    ProductsID      BIGINT REFERENCES Products(ID),
    ProductAddonsID BIGINT REFERENCES ProductAddonGroups(ID),
    Name           VARCHAR(255),
    Price          DOUBLE PRECISION
);

-- Insert Pizza and its addons
INSERT INTO Products (Name, Price) VALUES ('Pizza', 50000);
INSERT INTO ProductAddonGroups (ProductsID, Name, MaxQuantity) VALUES (1, 'Toppings', 1);
INSERT INTO ProductAddons (ProductsID, ProductAddonsID, Name, Price) VALUES 
  (1, 1, 'Cheese', 12000),
  (1, 1, 'Chicken', 18000),
  (1, 1, 'Pepper', 8000);

INSERT INTO ProductAddonGroups (ProductsID, Name, MaxQuantity) VALUES (1, 'Fillings', 1);
INSERT INTO ProductAddons (ProductsID, ProductAddonsID, Name, Price) VALUES 
  (1, 2, 'Cheese', 12000),
  (1, 2, 'Tomato', 9000),
  (1, 2, 'Tuna', 20000);

-- Insert Doughnut and its addons
INSERT INTO Products (Name, Price) VALUES ('Doughnut', 20000);
INSERT INTO ProductAddonGroups (ProductsID, Name, MaxQuantity) VALUES (2, 'Toppings', 1);
INSERT INTO ProductAddons (ProductsID, ProductAddonsID, Name, Price) VALUES 
  (2, 3, 'Blueberry', 12000),
  (2, 3, 'Cheese', 12000),
  (2, 3, 'Sugar Glaze', 10000);

INSERT INTO ProductAddonGroups (ProductsID, Name, MaxQuantity) VALUES (2, 'Fillings', 1);
INSERT INTO ProductAddons (ProductsID, ProductAddonsID, Name, Price) VALUES 
  (2, 4, 'Apple Slices', 14000),
  (2, 4, 'Milk Cream', 10000),
  (2, 4, 'Blueberry', 12000);

-- Insert Pie and its addons
INSERT INTO Products (Name, Price) VALUES ('Pie', 45000);
INSERT INTO ProductAddonGroups (ProductsID, Name, MaxQuantity) VALUES (3, 'Toppings', 1);
INSERT INTO ProductAddons (ProductsID, ProductAddonsID, Name, Price) VALUES 
  (3, 5, 'Pepper', 8000),
  (3, 5, 'Blueberry', 12000),
  (3, 5, 'Apple Slices', 14000);

INSERT INTO ProductAddonGroups (ProductsID, Name, MaxQuantity) VALUES (3, 'Fillings', 1);
INSERT INTO ProductAddons (ProductsID, ProductAddonsID, Name, Price) VALUES 
  (3, 6, 'Tuna', 20000),
  (3, 6, 'Cheese', 12000),
  (3, 6, 'Chicken', 18000);