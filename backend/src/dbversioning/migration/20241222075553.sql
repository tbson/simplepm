-- Modify "projects" table
ALTER TABLE "public"."projects" ADD CONSTRAINT "chk_projects_git_host" CHECK (git_host = ANY (ARRAY[''::text, 'GITHUB'::text, 'GITLAB'::text])), ADD COLUMN "git_repo" text NOT NULL DEFAULT '', ADD COLUMN "git_host" text NOT NULL DEFAULT '';
