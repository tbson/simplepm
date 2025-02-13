-- Create "git_accounts" table
CREATE TABLE "public"."git_accounts" (
  "id" bigserial NOT NULL,
  "tenant_id" bigint NULL,
  "uid" text NULL,
  "title" text NOT NULL DEFAULT '',
  "avatar" text NOT NULL DEFAULT '',
  "type" text NOT NULL DEFAULT 'GITHUB',
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_tenants_git_accounts" FOREIGN KEY ("tenant_id") REFERENCES "public"."tenants" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "chk_git_accounts_type" CHECK (type = ANY (ARRAY['GITHUB'::text, 'GITLAB'::text]))
);
-- Modify "projects" table
ALTER TABLE "public"."projects" DROP CONSTRAINT "chk_projects_git_host", DROP COLUMN "git_host", ADD COLUMN "git_account_id" bigint NULL, ADD CONSTRAINT "fk_projects_git_account" FOREIGN KEY ("git_account_id") REFERENCES "public"."git_accounts" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Drop "git_users" table
DROP TABLE "public"."git_users";
