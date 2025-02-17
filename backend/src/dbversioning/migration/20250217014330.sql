-- Create "git_pushes" table
CREATE TABLE "public"."git_pushes" (
  "id" bigserial NOT NULL,
  "task_id" bigint NULL,
  "user_id" bigint NULL,
  "git_account_uid" text NOT NULL,
  "git_repo" text NOT NULL,
  "git_host" text NOT NULL DEFAULT 'GITHUB',
  "git_branch" text NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_git_pushes_task" FOREIGN KEY ("task_id") REFERENCES "public"."tasks" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_git_pushes_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "chk_git_pushes_git_host" CHECK (git_host = ANY (ARRAY['GITHUB'::text, 'GITLAB'::text]))
);
-- Modify "git_commits" table
ALTER TABLE "public"."git_commits" DROP CONSTRAINT "chk_git_commits_git_host", DROP COLUMN "task_id", DROP COLUMN "user_id", DROP COLUMN "git_account_uid", DROP COLUMN "git_repo", DROP COLUMN "git_host", DROP COLUMN "git_branch", ADD COLUMN "git_push_id" bigint NULL, ADD CONSTRAINT "fk_git_pushes_git_commits" FOREIGN KEY ("git_push_id") REFERENCES "public"."git_pushes" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
