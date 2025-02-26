DROP TYPE IF EXISTS event.git_commit;
DROP TYPE IF EXISTS event.git_push;
DROP TABLE IF EXISTS event.attachments;
DROP TABLE IF EXISTS event.messages;

CREATE KEYSPACE IF NOT EXISTS event WITH replication = {
    'class': 'SimpleStrategy', 
    'replication_factor': 1
};

CREATE TYPE IF NOT EXISTS event.git_commit (
    id int,
    commit_id text,
    commit_url text,
    commit_message text,
    created_at timestamp,
);

CREATE TYPE IF NOT EXISTS event.git_push (
    id int,
    git_branch text,
    git_commits list<frozen<git_commit>>,
);

CREATE TABLE IF NOT EXISTS event.messages (
    id timeuuid,
    user_id int,
    task_id int,
    project_id int,
    content text,
    hash_tags list<text>,
    type text,
    git_push frozen<git_push>,
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
