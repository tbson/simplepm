-- Modify "projects" table
ALTER TABLE "public"."projects" ADD CONSTRAINT "chk_projects_git_host" CHECK (git_host = ANY (ARRAY['GITHUB'::text, 'GITLAB'::text])), DROP COLUMN "git_account_id", ADD COLUMN "git_host" text NOT NULL DEFAULT 'GITHUB';
-- Modify "users" table
ALTER TABLE "public"."users" DROP COLUMN "github_id", DROP COLUMN "gitlab_id";
-- Modify "git_branches" table
ALTER TABLE "public"."git_branches" ADD COLUMN "user_id" bigint NOT NULL, ADD COLUMN "created_at" timestamptz NULL, ADD COLUMN "updated_at" timestamptz NULL, ADD CONSTRAINT "fk_git_branches_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
