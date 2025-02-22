-- Create "changes" table
CREATE TABLE "public"."changes" (
  "id" bigserial NOT NULL,
  "tenant_id" bigint NOT NULL,
  "project_id" bigint NOT NULL,
  "task_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "user_full_name" text NOT NULL DEFAULT '',
  "source_type" text NOT NULL,
  "source_id" bigint NOT NULL,
  "source_title" text NOT NULL DEFAULT '',
  "action" text NOT NULL,
  "value" jsonb NOT NULL DEFAULT '{}',
  "created_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_changes_project" FOREIGN KEY ("project_id") REFERENCES "public"."projects" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_changes_task" FOREIGN KEY ("task_id") REFERENCES "public"."tasks" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_changes_tenant" FOREIGN KEY ("tenant_id") REFERENCES "public"."tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_changes_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
