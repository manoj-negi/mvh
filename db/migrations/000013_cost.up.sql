CREATE TABLE "cost" (
  "id" SERIAL PRIMARY KEY,
  "cost_input_id" int,
  "cost_type_id" int ,
  "period_id" int ,
  "title" varchar NOT NULL,
  "is_deleted" boolean DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  CONSTRAINT fk_cost_cost_input FOREIGN KEY (cost_input_id) REFERENCES cost_input(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_cost_cost_type FOREIGN KEY (cost_type_id) REFERENCES cost_type(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_cost_period_id FOREIGN KEY (period_id) REFERENCES period(id) ON UPDATE CASCADE ON DELETE RESTRICT
);



INSERT INTO cost (cost_input_id,cost_type_id,period_id,title) 
VALUES (1,1,1,'Parkbijdrage'),(null,1,1,'Schoonmaakkosten'),(1,3,1,'Verhuurbijjdrage'),
(1,2,1,'Electriciteit'),(1,2,1,'Overdrachtsbelasting'),(1,2,1,'Verzekering'),
(1,2,1,'Rioolbijdrage'),(1,2,1,'OZB-belasting'),(1,2,1,'Waterbelasting'),(2,3,1,'Rente'),
(2,1,1,'Klein onderhoud'),(1,2,4,'Groot onderhoud');