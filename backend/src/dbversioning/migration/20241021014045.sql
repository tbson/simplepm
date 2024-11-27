-- Create "auth_clients" table
CREATE TABLE "public"."auth_clients" (
  "id" bigserial NOT NULL,
  "uid" text NOT NULL,
  "description" text NOT NULL DEFAULT '',
  "secret" text NOT NULL,
  "partition" text NOT NULL,
  "default" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_auth_clients_uid" UNIQUE ("uid")
);
-- Create "variables" table
CREATE TABLE "public"."variables" (
  "id" bigserial NOT NULL,
  "key" text NOT NULL,
  "value" text NOT NULL DEFAULT '',
  "description" text NOT NULL DEFAULT '',
  "data_type" text NOT NULL DEFAULT 'STRING',
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_variables_key" UNIQUE ("key"),
  CONSTRAINT "chk_variables_data_type" CHECK (data_type = ANY (ARRAY['STRING'::text, 'INTEGER'::text, 'FLOAT'::text, 'BOOLEAN'::text, 'DATE'::text, 'DATETIME'::text]))
);
-- Create "tenants" table
CREATE TABLE "public"."tenants" (
  "id" bigserial NOT NULL,
  "auth_client_id" bigint NULL,
  "uid" text NOT NULL,
  "title" text NOT NULL,
  "avatar" text NOT NULL DEFAULT '',
  "avatar_str" text NOT NULL DEFAULT '',
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_tenants_uid" UNIQUE ("uid"),
  CONSTRAINT "fk_auth_clients_tenants" FOREIGN KEY ("auth_client_id") REFERENCES "public"."auth_clients" ("id") ON UPDATE NO ACTION ON DELETE SET NULL
);
-- Create "users" table
CREATE TABLE "public"."users" (
  "id" bigserial NOT NULL,
  "tenant_id" bigint NOT NULL,
  "tenant_tmp_id" bigint NULL,
  "uid" text NOT NULL,
  "email" text NOT NULL,
  "mobile" text NULL,
  "first_name" text NOT NULL DEFAULT '',
  "last_name" text NOT NULL DEFAULT '',
  "avatar" text NOT NULL DEFAULT '',
  "avatar_str" text NOT NULL DEFAULT '',
  "extra_info" jsonb NOT NULL DEFAULT '{}',
  "admin" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_users_tenant" FOREIGN KEY ("tenant_id") REFERENCES "public"."tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_users_tenant_email" to table: "users"
CREATE UNIQUE INDEX "idx_users_tenant_email" ON "public"."users" ("tenant_id", "email");
-- Create index "idx_users_tenant_uid" to table: "users"
CREATE UNIQUE INDEX "idx_users_tenant_uid" ON "public"."users" ("tenant_id", "uid");
