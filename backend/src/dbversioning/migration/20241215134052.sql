-- Modify "tasks" table
ALTER TABLE "public"."tasks" ADD
 CONSTRAINT "fk_features_tasks" FOREIGN KEY ("feature_id") REFERENCES "public"."features" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD
 CONSTRAINT "fk_projects_tasks" FOREIGN KEY ("project_id") REFERENCES "public"."projects" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
