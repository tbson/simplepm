-- Modify "task_fields" table
ALTER TABLE "public"."task_fields" ADD COLUMN "is_status" boolean NOT NULL DEFAULT false;
