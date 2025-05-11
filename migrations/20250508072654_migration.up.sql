create table users (
    id integer primary key,
    chat_id integer not null unique,
    created_at text not null,
    updated_at text not null
);