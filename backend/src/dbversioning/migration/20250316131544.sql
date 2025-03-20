-- Modify "tenants" table
ALTER TABLE "public"."tenants" DROP COLUMN "auth_client_id";
-- Drop "auth_clients" table
DROP TABLE "public"."auth_clients";
