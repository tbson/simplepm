-- Modify "docs" table
ALTER TABLE "public"."docs" DROP COLUMN "content";
ALTER TABLE "public"."docs" ADD COLUMN "content" jsonb NOT NULL DEFAULT '{}';
