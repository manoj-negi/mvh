CREATE TABLE "cost_type" (
  "id" SERIAL PRIMARY KEY,
  "description" varchar UNIQUE NOT NULL,
  "is_deleted" boolean DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

INSERT INTO cost_type (description) 
VALUES ('Verhuurkosten'),('Woningkosten'),('Financieringskosten'),
('Inkomstenbelasting');