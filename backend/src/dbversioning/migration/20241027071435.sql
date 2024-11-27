-- Create "roles" table
CREATE TABLE "public"."roles" (
  "id" bigserial NOT NULL,
  "tenant_id" bigint NOT NULL,
  "title" text NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_roles_tenant" FOREIGN KEY ("tenant_id") REFERENCES "public"."tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_roles_tenant_title" to table: "roles"
CREATE UNIQUE INDEX "idx_roles_tenant_title" ON "public"."roles" ("tenant_id", "title");
-- Create "pems" table
CREATE TABLE "public"."pems" (
  "id" bigserial NOT NULL,
  "title" text NOT NULL,
  "module" text NOT NULL,
  "action" text NOT NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_pems_module_action" to table: "pems"
CREATE UNIQUE INDEX "idx_pems_module_action" ON "public"."pems" ("module", "action");
-- Create "roles_pems" table
CREATE TABLE "public"."roles_pems" (
  "pem_id" bigint NOT NULL,
  "role_id" bigint NOT NULL,
  PRIMARY KEY ("pem_id", "role_id"),
  CONSTRAINT "fk_roles_pems_pem" FOREIGN KEY ("pem_id") REFERENCES "public"."pems" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_roles_pems_role" FOREIGN KEY ("role_id") REFERENCES "public"."roles" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "users_roles" table
CREATE TABLE "public"."users_roles" (
  "role_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  PRIMARY KEY ("role_id", "user_id"),
  CONSTRAINT "fk_users_roles_role" FOREIGN KEY ("role_id") REFERENCES "public"."roles" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_users_roles_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
