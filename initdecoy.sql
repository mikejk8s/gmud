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
CREATE DATABASE IF NOT EXISTS characters;
USE characters;
CREATE TABLE IF NOT EXISTS characters (
                                          id BIGINT UNIQUE NOT NULL PRIMARY KEY,
                                          name VARCHAR(30) UNIQUE NOT NULL,
    class VARCHAR(15) NOT NULL,
    race VARCHAR(15) NOT NULL DEFAULT 'HUMAN',
    level INT(3) NOT NULL DEFAULT '1',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    alive BOOLEAN NOT NULL DEFAULT '1',
    characterowner VARCHAR(20) NOT NULL DEFAULT 'player'
    );
-- Replace XXXXX with real variables and replace initdecoy.sql to init.sql when deploying.
CREATE USER 'XXXXXX'@'localhost' IDENTIFIED BY 'XXXXXXXX';
CREATE USER 'XXXXXXXX'@'%' IDENTIFIED BY 'XXXXXXXX';
CREATE USER 'XXXXXXXX'@'mysql' IDENTIFIED BY 'XXXXXXXX';
GRANT ALL PRIVILEGES ON *.* TO 'XXXXXXXX'@'localhost' WITH GRANT OPTION;
GRANT ALL PRIVILEGES ON *.* TO 'XXXXXXXX'@'%' WITH GRANT OPTION;
GRANT ALL PRIVILEGES ON *.* TO 'XXXXXXXX'@'mysql' WITH GRANT OPTION;
-- Make sure changes are applied.
FLUSH PRIVILEGES;