CREATE TABLE "cost_input" (
  "id" SERIAL PRIMARY KEY,
  "type" varchar UNIQUE NOT NULL,
  "is_deleted" boolean DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);


INSERT INTO cost_input (type) 
VALUES ('Amount'),('Percentage'),('Included in other amount');