-- Create "git_users" table
CREATE TABLE "public"."git_users" (
  "id" bigserial NOT NULL,
  "user_id" bigint NOT NULL,
  "username" text NOT NULL,
  "git_host" text NOT NULL DEFAULT 'GITHUB',
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_users_git_users" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "chk_git_users_git_host" CHECK (git_host = ANY (ARRAY['GITHUB'::text, 'GITLAB'::text]))
);
