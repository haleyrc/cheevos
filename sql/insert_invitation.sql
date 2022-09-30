INSERT INTO
	invitations (id, email, organization_id, expires_at, hashed_code)
VALUES
	($1, $2, $3, $4, $5);
