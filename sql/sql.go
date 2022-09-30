package sql

import _ "embed"

var (
	//go:embed delete_invitation.sql
	DeleteInvitationQuery string

	//go:embed get_cheevo.sql
	GetCheevoQuery string

	//go:embed get_invitation.sql
	GetInvitationQuery string

	//go:embed get_invitation_by_code.sql
	GetInvitationByCodeQuery string

	//go:embed get_membership.sql
	GetMembershipQuery string

	//go:embed get_user.sql
	GetUserQuery string

	//go:embed insert_award.sql
	InsertAwardQuery string

	//go:embed insert_cheevo.sql
	InsertCheevoQuery string

	//go:embed insert_invitation.sql
	InsertInvitationQuery string

	//go:embed insert_membership.sql
	InsertMembershipQuery string

	//go:embed insert_organization.sql
	InsertOrganizationQuery string

	//go:embed insert_user.sql
	InsertUserQuery string

	//go:embed update_invitation.sql
	UpdateInvitationQuery string
)
