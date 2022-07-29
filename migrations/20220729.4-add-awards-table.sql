CREATE TABLE awards (
  cheevo_id  UUID NOT NULL REFERENCES cheevos(id) ON DELETE CASCADE,
  awardee_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  awarder_id UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
	created_at TIMESTAMPTZ DEFAULT NOW(),
	PRIMARY KEY (cheevo_id, awardee_id)
);

COMMENT ON TABLE awards IS 'A cheevo that has been awarded to a user';
COMMENT ON COLUMN awards.cheevo_id IS 'The cheevo being awarded';
COMMENT ON COLUMN awards.awardee_id IS 'The user being awarded the cheevo';
COMMENT ON COLUMN awards.awarder_id IS 'The user that awarded the cheevo';
