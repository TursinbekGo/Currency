
CREATE TABLE "sale" (
    "id" UUID PRIMARY KEY,
    "user_id" UUID REFERENCES "users"("id"),
    "total_price" NUMERIC NOT NULL ,
    "total_count" BIGINT NOT NULL ,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);
