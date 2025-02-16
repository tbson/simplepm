-- Modify "git_commits" table
ALTER TABLE "public"."git_commits" ADD CONSTRAINT "chk_git_commits_git_host" CHECK (git_host = ANY (ARRAY['GITHUB'::text, 'GITLAB'::text])), ALTER COLUMN "task_id" DROP NOT NULL, DROP COLUMN "git_branch_id", ADD COLUMN "user_id" bigint NULL, ADD COLUMN "git_account_uid" text NOT NULL, ADD COLUMN "git_repo" text NOT NULL, ADD COLUMN "git_host" text NOT NULL DEFAULT 'GITHUB', ADD COLUMN "git_branch" text NOT NULL, ADD CONSTRAINT "fk_git_commits_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Drop "git_branches" table
DROP TABLE "public"."git_branches";
