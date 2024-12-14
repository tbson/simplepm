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
-- Create "workspaces" table
CREATE TABLE "public"."workspaces" (
  "id" bigserial NOT NULL,
  "tenant_id" bigint NOT NULL,
  "title" text NOT NULL,
  "description" text NOT NULL DEFAULT '',
  "avatar" text NOT NULL DEFAULT '',
  "order" bigint NOT NULL DEFAULT 0,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_workspaces_tenant" FOREIGN KEY ("tenant_id") REFERENCES "public"."tenants" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create "projects" table
CREATE TABLE "public"."projects" (
  "id" bigserial NOT NULL,
  "tenant_id" bigint NOT NULL,
  "workspace_id" bigint NULL,
  "title" text NOT NULL,
  "description" text NULL DEFAULT '',
  "avatar" text NOT NULL DEFAULT '',
  "layout" text NOT NULL DEFAULT 'TABLE',
  "status" text NOT NULL DEFAULT 'ACTIVE',
  "order" bigint NOT NULL DEFAULT 0,
  "finished_at" timestamptz NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_projects_tenant" FOREIGN KEY ("tenant_id") REFERENCES "public"."tenants" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "fk_projects_workspace" FOREIGN KEY ("workspace_id") REFERENCES "public"."workspaces" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "chk_projects_layout" CHECK (layout = ANY (ARRAY['TABLE'::text, 'KANBAN'::text, 'ROADMAP'::text])),
  CONSTRAINT "chk_projects_status" CHECK (status = ANY (ARRAY['ACTIVE'::text, 'ARCHIEVE'::text]))
);
-- Create "users" table
CREATE TABLE "public"."users" (
  "id" bigserial NOT NULL,
  "tenant_id" bigint NOT NULL,
  "tenant_tmp_id" bigint NULL,
  "sub" text NULL,
  "external_id" text NOT NULL,
  "email" text NOT NULL,
  "mobile" text NULL,
  "first_name" text NOT NULL DEFAULT '',
  "last_name" text NOT NULL DEFAULT '',
  "avatar" text NOT NULL DEFAULT '',
  "avatar_str" text NOT NULL DEFAULT '',
  "extra_info" jsonb NOT NULL DEFAULT '{}',
  "admin" boolean NOT NULL DEFAULT false,
  "locked_at" timestamp NULL,
  "locked_reason" text NOT NULL DEFAULT '',
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_users_sub" UNIQUE ("sub"),
  CONSTRAINT "fk_tenants_users" FOREIGN KEY ("tenant_id") REFERENCES "public"."tenants" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_users_tenant_email" to table: "users"
CREATE UNIQUE INDEX "idx_users_tenant_email" ON "public"."users" ("tenant_id", "email");
-- Create index "idx_users_tenant_external" to table: "users"
CREATE UNIQUE INDEX "idx_users_tenant_external" ON "public"."users" ("tenant_id", "external_id");
-- Create "projects_users" table
CREATE TABLE "public"."projects_users" (
  "id" bigserial NOT NULL,
  "project_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "creator_id" bigint NULL,
  "created_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_projects_users_creator" FOREIGN KEY ("creator_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL,
  CONSTRAINT "fk_projects_users_project" FOREIGN KEY ("project_id") REFERENCES "public"."projects" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "fk_projects_users_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_project_user" to table: "projects_users"
CREATE UNIQUE INDEX "idx_project_user" ON "public"."projects_users" ("project_id", "user_id");
-- Create "roles" table
CREATE TABLE "public"."roles" (
  "id" bigserial NOT NULL,
  "tenant_id" bigint NOT NULL,
  "title" text NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_tenants_roles" FOREIGN KEY ("tenant_id") REFERENCES "public"."tenants" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_roles_tenant_title" to table: "roles"
CREATE UNIQUE INDEX "idx_roles_tenant_title" ON "public"."roles" ("tenant_id", "title");
-- Create "pems" table
CREATE TABLE "public"."pems" (
  "id" bigserial NOT NULL,
  "title" text NOT NULL,
  "module" text NOT NULL,
  "action" text NOT NULL,
  "admin" boolean NOT NULL DEFAULT false,
  PRIMARY KEY ("id")
);
-- Create index "idx_pems_module_action" to table: "pems"
CREATE UNIQUE INDEX "idx_pems_module_action" ON "public"."pems" ("module", "action");
-- Create "roles_pems" table
CREATE TABLE "public"."roles_pems" (
  "pem_id" bigint NOT NULL,
  "role_id" bigint NOT NULL,
  PRIMARY KEY ("pem_id", "role_id"),
  CONSTRAINT "fk_roles_pems_pem" FOREIGN KEY ("pem_id") REFERENCES "public"."pems" ("id") ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT "fk_roles_pems_role" FOREIGN KEY ("role_id") REFERENCES "public"."roles" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);
-- Create "task_fields" table
CREATE TABLE "public"."task_fields" (
  "id" bigserial NOT NULL,
  "project_id" bigint NOT NULL,
  "title" text NOT NULL,
  "type" text NOT NULL,
  "description" text NOT NULL DEFAULT '',
  "order" bigint NOT NULL DEFAULT 0,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_task_fields_project" FOREIGN KEY ("project_id") REFERENCES "public"."projects" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create "task_field_options" table
CREATE TABLE "public"."task_field_options" (
  "id" bigserial NOT NULL,
  "task_field_id" bigint NOT NULL,
  "title" text NOT NULL,
  "description" text NULL DEFAULT '',
  "color" text NULL DEFAULT '',
  "order" bigint NOT NULL DEFAULT 0,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_task_fields_task_field_options" FOREIGN KEY ("task_field_id") REFERENCES "public"."task_fields" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "users_roles" table
CREATE TABLE "public"."users_roles" (
  "role_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  PRIMARY KEY ("role_id", "user_id"),
  CONSTRAINT "fk_users_roles_role" FOREIGN KEY ("role_id") REFERENCES "public"."roles" ("id") ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT "fk_users_roles_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);
-- Create "workspaces_users" table
CREATE TABLE "public"."workspaces_users" (
  "id" bigserial NOT NULL,
  "workspace_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "creator_id" bigint NULL,
  "created_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_workspaces_users_creator" FOREIGN KEY ("creator_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL,
  CONSTRAINT "fk_workspaces_users_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "fk_workspaces_users_workspace" FOREIGN KEY ("workspace_id") REFERENCES "public"."workspaces" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_workspace_user" to table: "workspaces_users"
CREATE UNIQUE INDEX "idx_workspace_user" ON "public"."workspaces_users" ("workspace_id", "user_id");
