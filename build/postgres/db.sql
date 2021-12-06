CREATE TABLE "products" (
  "id" SERIAL PRIMARY KEY,
  "sku" text,
  "name" text,
  "description" text,
  "category_id" int,
  "created_at" timestamp with time zone,
  "updated_at" timestamp with time zone
);

CREATE TABLE "categories" (
  "id" SERIAL PRIMARY KEY,
  "name" text,
  "description" text,
  "created_at" timestamp with time zone,
  "updated_at" timestamp with time zone
);

CREATE TABLE "stocks" (
  "product_id" int,
  "warehouse_id" int,
  "qty" double precision,
  "created_at" timestamp with time zone,
  "updated_at" timestamp with time zone
);

CREATE TABLE "product_prices" (
  "product_id" int,
  "original_price" double precision,
  "discount_percentage" double precision DEFAULT 1,
  "discount_amount" double precision DEFAULT 0,
  "expression" text,
  "created_at" timestamp with time zone,
  "updated_at" timestamp with time zone
);

CREATE TABLE "warehouses" (
  "id" SERIAL PRIMARY KEY,
  "code" text,
  "name" text
);

CREATE TABLE "orders" (
  "id" SERIAL PRIMARY KEY,
  "code" text,
  "customer_id" int,
  "status" text,
  "total_price" double precision,
  "created_at" timestamp with time zone,
  "updated_at" timestamp with time zone
);

CREATE TABLE "order_lines" (
  "id" SERIAL PRIMARY KEY,
  "product_id" int,
  "qty" int,
  "order_id" int,
  "price" double precision
);

CREATE TABLE "customers" (
  "id" SERIAL PRIMARY KEY,
  "name" text,
  "phone_number" text,
  "address" text,
  "ward" text,
  "district" text,
  "city" text,
  "country" text DEFAULT 'VN',
  "rank" text DEFAULT 'bronze',
  "created_at" timestamp with time zone,
  "updated_at" timestamp with time zone
);

CREATE TABLE "customer_ranks" (
  "id" SERIAL PRIMARY KEY,
  "code" text,
  "description" text,
  "expression" text
);

CREATE TABLE "users" (
    "id" serial PRIMARY KEY,
    "username" varchar,
    "password" varchar,
    "name" varchar
);

CREATE INDEX users_idx ON users ("username");

INSERT INTO users (username, password, name)
VALUES ('huy','$2a$04$o/d5Y8Phlg9o7T.CiohZ8ujUKsoE2fGd//vMG5FWArVy.4OaoQdeq','huy');
