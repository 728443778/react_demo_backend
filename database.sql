create database react_demo_backend with owner postgres;

\c react_demo_backend;

create table users (
    id serial not null primary key,
    username varchar(60) not null,
    password varchar(255) not null,
    status smallint not null,
    created_at timestamptz(6) not null,
    updated_at timestamptz(6) not null
);

create unique index unique_username on users using btree (username);    