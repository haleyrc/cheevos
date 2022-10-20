package repository

import _ "embed"

var (
	//go:embed sql/delete_invitation.sql
	DeleteInvitationQuery string

	//go:embed sql/get_cheevo.sql
	GetCheevoQuery string

	//go:embed sql/get_invitation.sql
	GetInvitationQuery string

	//go:embed sql/get_invitation_by_code.sql
	GetInvitationByCodeQuery string

	//go:embed sql/get_membership.sql
	GetMembershipQuery string

	//go:embed sql/get_user.sql
	GetUserQuery string

	//go:embed sql/insert_award.sql
	InsertAwardQuery string

	//go:embed sql/insert_cheevo.sql
	InsertCheevoQuery string

	//go:embed sql/insert_invitation.sql
	InsertInvitationQuery string

	//go:embed sql/insert_membership.sql
	InsertMembershipQuery string

	//go:embed sql/insert_organization.sql
	InsertOrganizationQuery string

	//go:embed sql/insert_user.sql
	InsertUserQuery string

	//go:embed sql/update_invitation.sql
	UpdateInvitationQuery string
)
