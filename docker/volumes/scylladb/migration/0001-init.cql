-- DROP TABLE IF EXISTS event.messages;
-- DROP TABLE IF EXISTS event.attachments;

CREATE KEYSPACE IF NOT EXISTS event WITH replication = {
    'class': 'SimpleStrategy', 
    'replication_factor': 1
};

CREATE TABLE IF NOT EXISTS event.messages (
    id timeuuid,
    user_id int,
    task_id int,
    project_id int,
    content text,
    hash_tags list<text>,
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
