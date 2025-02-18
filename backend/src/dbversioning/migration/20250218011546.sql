-- Modify "projects_users" table
ALTER TABLE "public"."projects_users" DROP COLUMN "creator_id";
-- Modify "tasks_users" table
ALTER TABLE "public"."tasks_users" DROP COLUMN "creator_id";
-- Modify "workspaces_users" table
ALTER TABLE "public"."workspaces_users" DROP COLUMN "creator_id";
