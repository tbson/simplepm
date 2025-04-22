-- Create "variables" table
CREATE TABLE "public"."variables" (
  "id" uuid NOT NULL DEFAULT public.uuid_generate_v1mc(),
  "key" text NOT NULL,
  "value" text NOT NULL DEFAULT '',
  "description" text NOT NULL DEFAULT '',
  "data_type" text NOT NULL DEFAULT 'STRING',
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_variables_key" UNIQUE ("key"),
  CONSTRAINT "chk_variables_data_type" CHECK (data_type = ANY (ARRAY['STRING'::text, 'INTEGER'::text, 'FLOAT'::text, 'BOOLEAN'::text, 'DATE'::text, 'DATETIME'::text]))
);
