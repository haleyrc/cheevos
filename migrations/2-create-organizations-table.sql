CREATE TABLE organizations (
  id       UUID PRIMARY KEY,
  name     TEXT NOT NULL,
  owner_id UUID NOT NULL REFERENCES users (id)
);
