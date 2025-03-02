DROP TABLE IF EXISTS event.messages;
DROP TABLE IF EXISTS event.attachments;
DROP TYPE IF EXISTS event.git_push;
DROP TYPE IF EXISTS event.git_commit;
DROP TYPE IF EXISTS event.git_pr;

CREATE KEYSPACE IF NOT EXISTS event WITH replication = {
    'class': 'SimpleStrategy', 
    'replication_factor': 1
};

CREATE TYPE IF NOT EXISTS event.git_commit (
    id text,
    commit_id text,
    commit_url text,
    commit_message text,
    created_at timestamp,
);

CREATE TYPE IF NOT EXISTS event.git_push (
    id text,
    git_branch text,
    git_commits list<frozen<git_commit>>,
);


CREATE TYPE IF NOT EXISTS event.git_pr (
    id text,
    title text,
    from_branch text,
    to_branch text,
    url text,
    merged_at timestamp,
    state text,
);


CREATE TABLE IF NOT EXISTS event.messages (
    id timeuuid,
    user_id bigint,
    task_id bigint,
    project_id bigint,
    content text,
    type text,
    git_push frozen<git_push>,
    git_pr frozen<git_pr>,
    user_name text,
    user_avatar text,
    user_color text,
    created_at timestamp,
    updated_at timestamp,
    PRIMARY KEY ((task_id), id)
);

CREATE TABLE IF NOT EXISTS event.attachments (
    id timeuuid,
    message_id timeuuid,
    file_name text,
    file_type text,
    file_url text,
    file_size int,
    thumbnail_url text,
    created_at timestamp,
    PRIMARY KEY ((message_id), id)
);
