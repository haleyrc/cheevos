UPDATE
	invitations
SET
	expires_at = $3,
	hashed_code = $4
WHERE
	email = $1
	AND
	organization_id = $2;
