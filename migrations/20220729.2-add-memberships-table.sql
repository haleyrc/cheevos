CREATE TABLE memberships (
  organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE RESTRICT,
  user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	created_at TIMESTAMPTZ DEFAULT NOW(),
	PRIMARY KEY (organization_id, user_id)
);

COMMENT ON TABLE memberships IS 'A connection between a user and an organization';
COMMENT ON COLUMN memberships.organization_id IS 'The organization the user is a member of';
COMMENT ON COLUMN memberships.user_id IS 'The user that is a member of the organization';
