-- Create "git_repos" table
CREATE TABLE "public"."git_repos" (
  "id" bigserial NOT NULL,
  "git_account_id" bigint NULL,
  "repo_id" text NOT NULL,
  "uid" text NOT NULL,
  "private" boolean NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_git_repos_git_account" FOREIGN KEY ("git_account_id") REFERENCES "public"."git_accounts" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
