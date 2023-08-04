
CREATE TABLE "sale_product" (
    "id" UUID PRIMARY KEY,
    "sale_id" UUID REFERENCES "sale"("id"),
    "product_id" UUID REFERENCES "product"("id"),
    "discount" BIGINT NOT NULL ,
    "discount_type" VARCHAR NOT NULL ,
    "product_name" VARCHAR NOT NULL ,
    "product_price" NUMERIC NOT NULL ,
    "price_with_discount" NUMERIC DEFAULT 0 ,
    "discount_price" NUMERIC DEFAULT 0 ,
    "count" BIGINT NOT NULL ,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);
