-- Modify "roles_pems" table
ALTER TABLE "public"."roles_pems" DROP CONSTRAINT "fk_roles_pems_pem", DROP CONSTRAINT "fk_roles_pems_role", ADD
 CONSTRAINT "fk_roles_pems_pem" FOREIGN KEY ("pem_id") REFERENCES "public"."pems" ("id") ON UPDATE CASCADE ON DELETE CASCADE, ADD
 CONSTRAINT "fk_roles_pems_role" FOREIGN KEY ("role_id") REFERENCES "public"."roles" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
-- Modify "users_roles" table
ALTER TABLE "public"."users_roles" DROP CONSTRAINT "fk_users_roles_role", DROP CONSTRAINT "fk_users_roles_user", ADD
 CONSTRAINT "fk_users_roles_role" FOREIGN KEY ("role_id") REFERENCES "public"."roles" ("id") ON UPDATE CASCADE ON DELETE CASCADE, ADD
 CONSTRAINT "fk_users_roles_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
