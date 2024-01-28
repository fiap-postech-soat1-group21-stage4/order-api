CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
DROP DATABASE IF EXISTS restaurant; 
CREATE DATABASE restaurant;
USE restaurant;

-- order
DROP TABLE IF EXISTS order;
CREATE TABLE order (
	ID  uuid DEFAULT uuid_generate_v4 (),
	CUSTOMER_ID  uuid DEFAULT uuid_generate_v4 (),
	STATUS VARCHAR(255) NOT NULL,
	CREATED_AT timestamptz NOT NULL DEFAULT now(),
	UPDATED_AT timestamptz NOT NULL DEFAULT now()
)

DROP TABLE IF EXISTS order_item;
CREATE TABLE order_item (
	ID  uuid DEFAULT uuid_generate_v4 (),
	ORDER_ID  uuid DEFAULT uuid_generate_v4 (),
	PRODUCT_ID  uuid DEFAULT uuid_generate_v4 (),
	QUANTITY DOUBLE PRECISION NOT NULL,
	CREATED_AT timestamptz NOT NULL DEFAULT now(),
	UPDATED_AT timestamptz NOT NULL DEFAULT now()
)