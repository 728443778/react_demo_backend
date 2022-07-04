create database react_demo_backend with owner postgres;

\c react_demo_backend;

create table users (
    id serial not null primary key,
    username varchar(60) not null,
    password varchar(255) not null,
    status int,
    created_at timestamptz(6),
    updated_at timestamptz(6)
);