-- Modify "roles" table
ALTER TABLE "public"."roles" DROP CONSTRAINT "fk_roles_tenant", ADD
 CONSTRAINT "fk_tenants_roles" FOREIGN KEY ("tenant_id") REFERENCES "public"."tenants" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
