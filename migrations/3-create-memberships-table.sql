CREATE TABLE memberships (
	organization_id UUID        NOT NULL REFERENCES organizations(id),
	user_id         UUID        NOT NULL REFERENCES users(id),
	joined_at       TIMESTAMPTZ NOT NULL,
	PRIMARY KEY (organization_id, user_id)
);

CREATE INDEX idx_memberships_organization_id ON memberships(organization_id);
CREATE INDEX idx_memberships_user_id ON memberships(user_id);

