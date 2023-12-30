CREATE TABLE "period" (
  "id" SERIAL PRIMARY KEY,
  "period" varchar UNIQUE NOT NULL,
  "is_deleted" boolean DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

INSERT INTO period (period) 
VALUES ('Year'),('Quarter'),('Month'),('10_Year');