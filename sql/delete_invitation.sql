DELETE FROM
	invitations
WHERE
	hashed_code = $1;
