CREATE DATABASE IF NOT EXISTS users;
USE users;
CREATE TABLE IF NOT EXISTS users(
    id            integer unsigned null,
    created_at    datetime         null,
    updated_at    datetime         null,
    deleted_at    datetime         null,
    name          varchar(255)     null,
    password_hash varchar(255)     null,
    remember_hash varchar(255)     null
);
-- Replace XXXXX with real variables and replace initdecoy.sql to init.sql when deploying.
CREATE USER 'XXXXXX'@'localhost' IDENTIFIED WITH mysql_native_password BY 'XXXXXXXX';
CREATE USER 'XXXXXXXX'@'%' IDENTIFIED WITH mysql_native_password BY 'XXXXXXXX';
CREATE USER 'XXXXXXXX'@'mysql' IDENTIFIED WITH mysql_native_password BY 'XXXXXXXX';
GRANT ALL PRIVILEGES ON *.* TO 'XXXXXXXX'@'localhost' WITH GRANT OPTION;
GRANT ALL PRIVILEGES ON *.* TO 'XXXXXXXX'@'%' WITH GRANT OPTION;
GRANT ALL PRIVILEGES ON *.* TO 'XXXXXXXX'@'mysql' WITH GRANT OPTION;
-- Make sure changes are applied.
FLUSH PRIVILEGES;