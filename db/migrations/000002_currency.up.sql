CREATE TABLE "currency" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "symbol" varchar(10) NOT NULL,
  "isdeleted" boolean DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);