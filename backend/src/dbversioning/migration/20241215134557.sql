-- Modify "task_fields" table
ALTER TABLE "public"."task_fields" ADD CONSTRAINT "chk_task_fields_type" CHECK (type = ANY (ARRAY['TEXT'::text, 'NUMBER'::text, 'DATE'::text, 'SELECT'::text, 'MULTIPLE_SELECT'::text])), ALTER COLUMN "type" SET DEFAULT 'TEXT';
