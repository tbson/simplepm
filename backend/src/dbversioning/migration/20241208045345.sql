-- Modify "projects" table
ALTER TABLE "public"."projects" ADD CONSTRAINT "chk_projects_status" CHECK (status = ANY (ARRAY['ACTIVE'::text, 'ARCHIEVE'::text])), DROP COLUMN "start_date", DROP COLUMN "target_date", ADD COLUMN "status" text NOT NULL DEFAULT 'ACTIVE';
