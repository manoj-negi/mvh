CREATE TABLE "permission" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "permission" varchar UNIQUE NOT NULL,
  "is_deleted" boolean DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);