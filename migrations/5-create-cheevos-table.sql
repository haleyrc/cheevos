CREATE TABLE cheevos (
	id              UUID PRIMARY KEY,
	organization_id UUID NOT NULL REFERENCES organizations(id),
	name            TEXT NOT NULL,
	description     TEXT NOT NULL
);

CREATE INDEX idx_cheevos_organization_id ON cheevos(id);
