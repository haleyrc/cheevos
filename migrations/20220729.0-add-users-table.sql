CREATE EXTENSION pgcrypto;

CREATE TABLE users (
	id         UUID PRIMARY KEY,
	username   TEXT UNIQUE NOT NULL CHECK (length(username) > 0),
	created_at TIMESTAMPTZ DEFAULT NOW()
);

COMMENT ON TABLE users IS 'Users of the application';
COMMENT ON COLUMN users.id IS 'A unique identifier for the user';
COMMENT ON COLUMN users.username IS 'A unique user-facing identifier for the user';
