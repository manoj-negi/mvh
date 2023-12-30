CREATE TABLE "entry_fields_cost" (
  "id" SERIAL PRIMARY KEY,
  "cost_id" int NOT NULL,
  "entry_field_id" int NOT NULL,
  "is_deleted" boolean DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  CONSTRAINT fk_cost_cost_input FOREIGN KEY (cost_id) REFERENCES cost(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_cost_cost_type FOREIGN KEY (entry_field_id) REFERENCES entry_fields(id) ON UPDATE CASCADE ON DELETE RESTRICT
);
