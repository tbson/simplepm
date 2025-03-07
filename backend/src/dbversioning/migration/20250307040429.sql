-- Modify "users" table
ALTER TABLE "public"."users" ADD COLUMN "pwd" text NOT NULL DEFAULT '', ADD COLUMN "reset_pwd_token" text NOT NULL DEFAULT '', ADD COLUMN "reset_pwd_expiry" timestamp NULL;
