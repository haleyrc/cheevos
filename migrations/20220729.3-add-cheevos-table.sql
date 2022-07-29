CREATE TABLE cheevos (
	id              UUID PRIMARY KEY,
	name            TEXT NOT NULL CHECK(length(name) > 0),
	description     TEXT NOT NULL CHECK(length(description) > 0),
  organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
	created_at      TIMESTAMPTZ DEFAULT NOW()
);

COMMENT ON TABLE cheevos IS 'An achievement that can awarded to users of an organization';
COMMENT ON COLUMN cheevos.name IS 'The short name for the cheevo';
COMMENT ON COLUMN cheevos.description IS 'The long form description of the cheevo';
COMMENT ON COLUMN cheevos.organization_id IS 'The oragnization the cheevo belongs to';
