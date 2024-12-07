-- Modify "task_field_options" table
ALTER TABLE "public"."task_field_options" ADD COLUMN "description" text NULL DEFAULT '', ADD COLUMN "color" text NULL DEFAULT '';
