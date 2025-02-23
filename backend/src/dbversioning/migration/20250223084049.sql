-- Modify "changes" table
ALTER TABLE "public"."changes" DROP CONSTRAINT "fk_changes_project", DROP CONSTRAINT "fk_changes_task", DROP CONSTRAINT "fk_changes_tenant", DROP CONSTRAINT "fk_changes_user";
