-- Modify "users" table
ALTER TABLE "public"."users" DROP CONSTRAINT "fk_users_tenant", ADD
 CONSTRAINT "fk_tenants_users" FOREIGN KEY ("tenant_id") REFERENCES "public"."tenants" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
