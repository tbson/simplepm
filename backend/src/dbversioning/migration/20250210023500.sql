-- Modify "docs" table
ALTER TABLE "public"."docs" ADD COLUMN "file_name" text NOT NULL DEFAULT '', ADD COLUMN "file_type" text NOT NULL DEFAULT '', ADD COLUMN "file_size" bigint NOT NULL DEFAULT 0, ADD COLUMN "file_url" text NOT NULL DEFAULT '';
-- Create "doc_attachments" table
CREATE TABLE "public"."doc_attachments" (
  "id" bigserial NOT NULL,
  "user_id" bigint NOT NULL,
  "task_id" bigint NOT NULL,
  "file_name" text NOT NULL,
  "file_type" text NOT NULL,
  "file_size" bigint NOT NULL,
  "file_url" text NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_doc_attachments_task" FOREIGN KEY ("task_id") REFERENCES "public"."tasks" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_doc_attachments_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
