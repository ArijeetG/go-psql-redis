drop database if exists test;
create database test;
\c test;
create table "user"(
    user_id serial not null primary key,
    name varchar(256) not null,
    password varchar(256) not null,
    user_type varchar(256) not null,
    address varchar(256) default 'null',
    phone varchar(256) default 'null'
);