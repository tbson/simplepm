-- Rename a column from "reset_pwd_token" to "pwd_reset_token"
ALTER TABLE "public"."users" RENAME COLUMN "reset_pwd_token" TO "pwd_reset_token";
-- Rename a column from "reset_pwd_expiry" to "pwd_reset_at"
ALTER TABLE "public"."users" RENAME COLUMN "reset_pwd_expiry" TO "pwd_reset_at";
