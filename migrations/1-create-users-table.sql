CREATE EXTENSION pgcrypto;

CREATE TABLE users (
  id            UUID PRIMARY KEY,
  username      TEXT UNIQUE NOT NULL,
  password_hash TEXT NOT NULL
);

CREATE INDEX idx_users_username ON users (username);
