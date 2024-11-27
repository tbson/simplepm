-- Drop index "idx_users_tenant_uid" from table: "users"
DROP INDEX "public"."idx_users_tenant_uid";
-- Rename a column from "uid" to "external_id"
ALTER TABLE "public"."users" RENAME COLUMN "uid" TO "external_id";
-- Modify "users" table
ALTER TABLE "public"."users" ADD COLUMN "sub" text NULL, ADD CONSTRAINT "uni_users_sub" UNIQUE ("sub");
-- Create index "idx_users_tenant_external" to table: "users"
CREATE UNIQUE INDEX "idx_users_tenant_external" ON "public"."users" ("tenant_id", "external_id");
