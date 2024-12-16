-- Modify "features" table
ALTER TABLE "public"."features" ADD COLUMN "default" boolean NOT NULL DEFAULT false;
