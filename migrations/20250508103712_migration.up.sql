create table applications (
    id integer primary key,
    user_id integer not null,
    country varchar(255) null,
    mark_or_conditions varchar(255) null,
    budget varchar(255) null,
    steering_wheel_type varchar(255) null,
    city varchar(255) null,
    person_name varchar(255) null,
    person_phone varchar(255) null,
    submitted boolean null,
    step integer not null default 1,
    created_at text not null,
    updated_at text not null
);
