SELECT
	organization_id, user_id, joined_at
FROM
	memberships
WHERE
	organization_id = $1
	AND
	user_id = $2;
