-- Create "docs" table
CREATE TABLE "public"."docs" (
  "id" bigserial NOT NULL,
  "user_id" bigint NOT NULL,
  "task_id" bigint NOT NULL,
  "type" text NOT NULL DEFAULT 'DOC',
  "title" text NOT NULL,
  "description" text NOT NULL DEFAULT '',
  "content" text NOT NULL DEFAULT '',
  "url" text NOT NULL DEFAULT '',
  "order" bigint NOT NULL DEFAULT 0,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_docs_task" FOREIGN KEY ("task_id") REFERENCES "public"."tasks" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_docs_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "chk_docs_type" CHECK (type = ANY (ARRAY['DOC'::text, 'FILE'::text, 'LINK'::text]))
);
