SELECT
	email, organization_id, expires_at
FROM
	invitations
WHERE
	id = $1;
