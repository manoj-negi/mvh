CREATE TABLE "static_values" (
  "id" SERIAL PRIMARY KEY,
  "cost_input_id" int NOT NULL,
  "description" varchar UNIQUE NOT NULL,
  "value" float NOT NULL,
  "is_percentage" boolean DEFAULT false,
  "is_deleted" boolean DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  CONSTRAINT fk_static_values_cost_input FOREIGN KEY (cost_input_id) REFERENCES cost_input(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

INSERT INTO static_values (cost_input_id,description,value,is_percentage) 
VALUES (2,'Overdrachtsbelasting',10.30,true),(2,'Box 3_forfaitair rendement',6.17,true),
(2,'Box 3_berekening_sparen',0.30,true),(2,'Box 3_vermogensbelasting',32,true),
(2,'Woz % woning',60,true),(2,'Inflatie per jaar',2,true),
(2,'Spaarrrente',3,true),(1,'Notaris kosten',750,false),
(1,'Verbouwtermijn',10,false);