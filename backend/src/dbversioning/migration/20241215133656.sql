-- Create "features" table
CREATE TABLE "public"."features" (
  "id" bigserial NOT NULL,
  "project_id" bigint NOT NULL,
  "title" text NOT NULL,
  "description" text NULL DEFAULT '',
  "status" text NOT NULL DEFAULT 'ACTIVE',
  "order" bigint NOT NULL DEFAULT 0,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_features_project" FOREIGN KEY ("project_id") REFERENCES "public"."tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "chk_features_status" CHECK (status = ANY (ARRAY['ACTIVE'::text, 'FINISHED'::text, 'ARCHIEVED'::text]))
);
-- Create "tasks" table
CREATE TABLE "public"."tasks" (
  "id" bigserial NOT NULL,
  "project_id" bigint NOT NULL,
  "feature_id" bigint NOT NULL,
  "user_id" bigint NULL,
  "title" text NOT NULL,
  "description" text NULL DEFAULT '',
  "order" bigint NOT NULL DEFAULT 0,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_tasks_feature" FOREIGN KEY ("feature_id") REFERENCES "public"."tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_tasks_project" FOREIGN KEY ("project_id") REFERENCES "public"."tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_tasks_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "task_field_values" table
CREATE TABLE "public"."task_field_values" (
  "id" bigserial NOT NULL,
  "task_id" bigint NOT NULL,
  "task_field_id" bigint NOT NULL,
  "task_field_option_id" bigint NULL,
  "number_value" bigint NULL,
  "date_value" timestamptz NULL,
  "value" text NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_task_field_options_task_field_values" FOREIGN KEY ("task_field_option_id") REFERENCES "public"."task_field_options" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "fk_task_fields_task_field_values" FOREIGN KEY ("task_field_id") REFERENCES "public"."task_fields" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "fk_tasks_task_field_values" FOREIGN KEY ("task_id") REFERENCES "public"."tasks" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
