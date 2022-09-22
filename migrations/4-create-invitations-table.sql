CREATE TABLE invitations (
	email           TEXT        NOT NULL,
	organization_id UUID        NOT NULL REFERENCES organizations(id),
	expires_at      TIMESTAMPTZ NOT NULL,
	hashed_code     TEXT        NOT NULL,
	PRIMARY KEY (email, organization_id)
);

CREATE INDEX idx_invitations_email ON invitations(email);
CREATE INDEX idx_invitations_organization_id ON invitations(organization_id);
