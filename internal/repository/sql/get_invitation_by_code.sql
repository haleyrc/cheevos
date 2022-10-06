SELECT
	email, organization_id, expires_at
FROM
	invitations
WHERE
	hashed_code = $1;
