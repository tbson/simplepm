-- Rename a column from "git_repo" to "git_repo_uid"
ALTER TABLE "public"."git_pushes" RENAME COLUMN "git_repo" TO "git_repo_uid";
