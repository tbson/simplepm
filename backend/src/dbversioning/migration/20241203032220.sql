-- Modify "workspaces_users" table
ALTER TABLE "public"."workspaces_users" ALTER COLUMN "creator_id" DROP NOT NULL;
-- Create "projects" table
CREATE TABLE "public"."projects" (
  "id" bigserial NOT NULL,
  "tenant_id" bigint NOT NULL,
  "workspace_id" bigint NULL,
  "title" text NOT NULL,
  "description" text NULL DEFAULT '',
  "avatar" text NOT NULL DEFAULT '',
  "layout" text NOT NULL DEFAULT 'TABLE',
  "order" bigint NOT NULL DEFAULT 0,
  "start_date" timestamptz NULL,
  "target_date" timestamptz NULL,
  "finished_at" timestamptz NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_projects_tenant" FOREIGN KEY ("tenant_id") REFERENCES "public"."tenants" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "fk_projects_workspace" FOREIGN KEY ("workspace_id") REFERENCES "public"."workspaces" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "chk_projects_layout" CHECK (layout = ANY (ARRAY['TABLE'::text, 'KANBAN'::text, 'ROADMAP'::text]))
);
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
  "order" bigint NOT NULL DEFAULT 0,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_task_field_options_task_field" FOREIGN KEY ("task_field_id") REFERENCES "public"."task_fields" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
