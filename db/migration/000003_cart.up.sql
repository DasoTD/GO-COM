CREATE TABLE "carts" (
    "owner" varchar NOT NULL,
    "product" varchar NOT NULL,
    "quantity" bigint NOT NULL
);

ALTER TABLE "carts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");
-- ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");
