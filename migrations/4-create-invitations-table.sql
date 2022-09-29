CREATE TABLE invitations (
	id              UUID        PRIMARY KEY,
	email           TEXT        NOT NULL,
	organization_id UUID        NOT NULL REFERENCES organizations(id),
	expires_at      TIMESTAMPTZ NOT NULL,
	hashed_code     TEXT        NOT NULL
);

CREATE INDEX idx_invitations_hashed_code ON invitations(hashed_code);
CREATE INDEX idx_invitations_organization_id ON invitations(organization_id);
