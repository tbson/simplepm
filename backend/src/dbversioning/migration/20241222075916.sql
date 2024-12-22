-- Create "git_branches" table
CREATE TABLE "public"."git_branches" (
  "id" bigserial NOT NULL,
  "task_id" bigint NOT NULL,
  "title" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_tasks_git_branches" FOREIGN KEY ("task_id") REFERENCES "public"."tasks" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
