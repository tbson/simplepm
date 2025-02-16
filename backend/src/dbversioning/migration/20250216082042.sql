-- Create "git_commits" table
CREATE TABLE "public"."git_commits" (
  "id" bigserial NOT NULL,
  "task_id" bigint NOT NULL,
  "git_branch_id" bigint NOT NULL,
  "commit_id" text NOT NULL,
  "commit_url" text NOT NULL,
  "commit_message" text NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_git_branches_git_commits" FOREIGN KEY ("git_branch_id") REFERENCES "public"."git_branches" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "fk_git_commits_task" FOREIGN KEY ("task_id") REFERENCES "public"."tasks" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
