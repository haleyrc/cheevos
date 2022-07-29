CREATE TABLE organizations (
	id         UUID PRIMARY KEY,
	name       TEXT NOT NULL CHECK (length(name) > 0),
  owner      UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
	created_at TIMESTAMPTZ DEFAULT NOW()
);

COMMENT ON TABLE organizations IS 'A group of users that can award cheevos between them';
COMMENT ON COLUMN organizations.id IS 'A unique identifier for the organization';
COMMENT ON COLUMN organizations.name IS 'The name for the organization';
