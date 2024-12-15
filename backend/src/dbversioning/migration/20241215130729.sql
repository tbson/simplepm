-- Modify "workspaces" table
ALTER TABLE "public"."workspaces" DROP CONSTRAINT "fk_workspaces_tenant", ADD
 CONSTRAINT "fk_workspaces_tenant" FOREIGN KEY ("tenant_id") REFERENCES "public"."tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "projects" table
ALTER TABLE "public"."projects" DROP CONSTRAINT "fk_projects_tenant", DROP CONSTRAINT "fk_projects_workspace", DROP CONSTRAINT "chk_projects_status", ADD CONSTRAINT "chk_projects_status" CHECK (status = ANY (ARRAY['ACTIVE'::text, 'FINISHED'::text, 'ARCHIEVE'::text])), ADD
 CONSTRAINT "fk_projects_tenant" FOREIGN KEY ("tenant_id") REFERENCES "public"."tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD
 CONSTRAINT "fk_workspaces_projects" FOREIGN KEY ("workspace_id") REFERENCES "public"."workspaces" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "projects_users" table
ALTER TABLE "public"."projects_users" DROP CONSTRAINT "fk_projects_users_creator", DROP CONSTRAINT "fk_projects_users_project", DROP CONSTRAINT "fk_projects_users_user", ADD
 CONSTRAINT "fk_projects_users_creator" FOREIGN KEY ("creator_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD
 CONSTRAINT "fk_projects_users_project" FOREIGN KEY ("project_id") REFERENCES "public"."projects" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD
 CONSTRAINT "fk_projects_users_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "task_fields" table
ALTER TABLE "public"."task_fields" DROP CONSTRAINT "fk_task_fields_project", ADD
 CONSTRAINT "fk_projects_task_fields" FOREIGN KEY ("project_id") REFERENCES "public"."projects" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "task_field_options" table
ALTER TABLE "public"."task_field_options" DROP CONSTRAINT "fk_task_fields_task_field_options", ADD
 CONSTRAINT "fk_task_fields_task_field_options" FOREIGN KEY ("task_field_id") REFERENCES "public"."task_fields" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "workspaces_users" table
ALTER TABLE "public"."workspaces_users" DROP CONSTRAINT "fk_workspaces_users_creator", DROP CONSTRAINT "fk_workspaces_users_user", DROP CONSTRAINT "fk_workspaces_users_workspace", ADD
 CONSTRAINT "fk_workspaces_users_creator" FOREIGN KEY ("creator_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD
 CONSTRAINT "fk_workspaces_users_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD
 CONSTRAINT "fk_workspaces_users_workspace" FOREIGN KEY ("workspace_id") REFERENCES "public"."workspaces" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
