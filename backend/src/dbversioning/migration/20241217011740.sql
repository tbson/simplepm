-- Create "features_users" table
CREATE TABLE "public"."features_users" (
  "feature_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  PRIMARY KEY ("feature_id", "user_id"),
  CONSTRAINT "fk_features_users_feature" FOREIGN KEY ("feature_id") REFERENCES "public"."features" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_features_users_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Modify "tasks" table
ALTER TABLE "public"."tasks" DROP COLUMN "user_id";
-- Create "tasks_users" table
CREATE TABLE "public"."tasks_users" (
  "task_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  PRIMARY KEY ("task_id", "user_id"),
  CONSTRAINT "fk_tasks_users_task" FOREIGN KEY ("task_id") REFERENCES "public"."tasks" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_tasks_users_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
