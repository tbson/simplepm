-- Modify "tasks_users" table
ALTER TABLE "public"."tasks_users" DROP CONSTRAINT "fk_tasks_users_task", ADD CONSTRAINT "fk_tasks_task_users" FOREIGN KEY ("task_id") REFERENCES "public"."tasks" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
