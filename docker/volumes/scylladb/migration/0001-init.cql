CREATE KEYSPACE IF NOT EXISTS event WITH replication = {
    'class': 'SimpleStrategy', 
    'replication_factor': 1
};

CREATE TABLE IF NOT EXISTS event.messages (
    user_id int,
    id timeuuid,
    task_id int,
    project_id int,
    workspace_id int,
    tenant_id int,
    type int,
    content text,
    tags list<text>,
    doc_id int,
    doc_title text,
    doc_description text,
    doc_thumbnail text,
    doc_url text,
    created_at timestamp,
    updated_at timestamp,
    PRIMARY KEY ((task_id), id)
);
