CREATE TABLE "brand" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "logo" varchar DEFAULT NULL,
  "website" varchar DEFAULT NULL,
  "validated" boolean DEFAULT false,
  "is_deleted" boolean DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

INSERT INTO brand (name,website,validated) 
VALUES 
('Landal GreenParks','http://www.landal.nl',true),
('Roompot','http://www.roompot.nl',true),
('Center Parcs','http://www.centerparcs.nl',true),
('Topparken','http://www.topparken.nl',true),
('Europarcs','http://www.europarcs.nl',true),
('Dutchen','http://www.dutchen.nl',true),
('Uplandparcs','http://www.uplandparcs.nl',true),
('Dormio','http://www.dormio.nl',true),
('Summio','http://www.summioparcs.nl',true),
('Andere','',true);
