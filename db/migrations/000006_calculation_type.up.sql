CREATE TABLE "calculation_type" (
  "id" SERIAL PRIMARY KEY,
  "description" varchar UNIQUE NOT NULL,
  "is_deleted" boolean DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

INSERT INTO calculation_type (description) 
VALUES ('Recreatiewoning'),('Huurwoning'),('Camper');