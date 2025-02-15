-- Modify "users" table
ALTER TABLE "public"."users" ADD COLUMN "github_id" text NOT NULL DEFAULT '', ADD COLUMN "gitlab_id" text NOT NULL DEFAULT '';
