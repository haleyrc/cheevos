SELECT
	id, name, description, organization_id
FROM
	cheevos
WHERE
	id = $1;
