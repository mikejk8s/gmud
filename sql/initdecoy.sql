CREATE DATABASE IF NOT EXISTS users;

\c users;

CREATE TABLE IF NOT EXISTS users (
    id            SERIAL PRIMARY KEY,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP,
    deleted_at    TIMESTAMP,
    name          VARCHAR(255),
    password_hash VARCHAR(255),
    remember_hash VARCHAR(255)
);

CREATE DATABASE IF NOT EXISTS characters;

\c characters;

CREATE TABLE IF NOT EXISTS characters (
    id              BIGSERIAL PRIMARY KEY,
    name            VARCHAR(30) UNIQUE NOT NULL,
    class           VARCHAR(15) NOT NULL,
    race            VARCHAR(15) NOT NULL DEFAULT 'HUMAN',
    level           INTEGER NOT NULL DEFAULT '1',
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    alive           BOOLEAN NOT NULL DEFAULT '1',
    characterowner  VARCHAR(20) NOT NULL DEFAULT 'player'
);

CREATE USER gmud WITH PASSWORD 'gmud';

GRANT ALL PRIVILEGES ON DATABASE users TO gmud;
GRANT ALL PRIVILEGES ON DATABASE characters TO gmud;