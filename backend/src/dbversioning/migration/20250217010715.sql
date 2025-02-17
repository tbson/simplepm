-- Modify "tasks_users" table
ALTER TABLE "public"."tasks_users" DROP CONSTRAINT "tasks_users_pkey", ADD COLUMN "id" bigserial NOT NULL, ADD COLUMN "creator_id" bigint NULL, ADD COLUMN "git_branch" text NULL, ADD COLUMN "created_at" timestamptz NULL, ADD COLUMN "update_at" timestamptz NULL, ADD PRIMARY KEY ("id"), ADD CONSTRAINT "fk_tasks_users_creator" FOREIGN KEY ("creator_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Create index "idx_task_user_branch" to table: "tasks_users"
CREATE UNIQUE INDEX "idx_task_user_branch" ON "public"."tasks_users" ("task_id", "user_id", "git_branch");
