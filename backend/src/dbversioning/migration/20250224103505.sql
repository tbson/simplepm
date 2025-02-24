-- Modify "users" table
ALTER TABLE "public"."users" ADD COLUMN "github_username" text NOT NULL DEFAULT '';
