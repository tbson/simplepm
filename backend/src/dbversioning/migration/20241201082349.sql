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
-- Create "workspaces_users" table
CREATE TABLE "public"."workspaces_users" (
  "id" bigserial NOT NULL,
  "workspace_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "creator_id" bigint NOT NULL,
  "created_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_workspaces_users_creator" FOREIGN KEY ("creator_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL,
  CONSTRAINT "fk_workspaces_users_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "fk_workspaces_users_workspace" FOREIGN KEY ("workspace_id") REFERENCES "public"."workspaces" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_workspace_user" to table: "workspaces_users"
CREATE UNIQUE INDEX "idx_workspace_user" ON "public"."workspaces_users" ("workspace_id", "user_id");
