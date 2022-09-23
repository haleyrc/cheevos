CREATE TABLE awards (
	cheevo_id  UUID        NOT NULL REFERENCES cheevos(id),
	user_id    UUID        NOT NULL REFERENCES users(id),
	awarded_at TIMESTAMPTZ NOT NULL,
	PRIMARY KEY (cheevo_id, user_id)
);

CREATE INDEX idx_awards_cheevo_id ON awards(cheevo_id);
CREATE INDEX idx_awards_user_id ON awards(user_id);
