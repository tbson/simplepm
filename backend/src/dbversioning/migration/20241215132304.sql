-- Modify "projects" table
ALTER TABLE "public"."projects" DROP CONSTRAINT "chk_projects_status", ADD CONSTRAINT "chk_projects_status" CHECK (status = ANY (ARRAY['ACTIVE'::text, 'FINISHED'::text, 'ARCHIEVED'::text]));
