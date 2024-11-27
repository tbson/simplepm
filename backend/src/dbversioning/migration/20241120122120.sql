-- Modify "users" table
ALTER TABLE "public"."users" ADD COLUMN "locked_reason" text NOT NULL DEFAULT '';
