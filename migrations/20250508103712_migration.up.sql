create table applications (
    id integer primary key,
    chat_id integer not null,
    telegram_id varchar(255) not null,
    user_id integer not null,
    country varchar(255) null,
    mark_or_conditions varchar(255) null,
    budget varchar(255) null,
    steering_wheel_type varchar(255) null,
    city varchar(255) null,
    person_name varchar(255) null,
    person_phone varchar(255) null,
    step integer not null default 1,
    created_at text not null,
    updated_at text not null,
    sended_telegram boolean default false,
    sended_bitrix boolean default false
);
