CREATE TABLE "roles_permission" (
  "id" SERIAL PRIMARY KEY,
  "role_id" int NOT NULL,
  "permission_id" int NOT NULL,
  "is_deleted" boolean DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  CONSTRAINT fk_role_permission_role FOREIGN KEY (role_id) REFERENCES role(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_permission_permission_role FOREIGN KEY (permission_id) REFERENCES permission(id) ON UPDATE CASCADE ON DELETE RESTRICT
);



