package server

import (
	"context"

	"github.com/haleyrc/cheevos/auth"
	"github.com/haleyrc/cheevos/cheevos"
	"github.com/haleyrc/cheevos/roster"
)

type AuthorizationService interface {
	CanAwardCheevo(ctx context.Context, awarderID, recipientID, cheevoID string) error
	CanCreateCheevo(ctx context.Context, userID, orgID string) error
	CanGetCheevo(ctx context.Context, userID, cheevoID string) error
	CanInviteUsersToOrganization(ctx context.Context, userID, orgID string) error
	CanRefreshInvitation(ctx context.Context, userID, invitationID string) error
}

type AuthenticationService interface {
	SignUp(ctx context.Context, username, password string) (*auth.User, error)
}

type CheevosService interface {
	AwardCheevoToUser(ctx context.Context, recipientID, cheevoID string) error
	CreateCheevo(ctx context.Context, name, description, orgID string) (*cheevos.Cheevo, error)
	GetCheevo(ctx context.Context, id string) (*cheevos.Cheevo, error)
}

type RosterService interface {
	AcceptInvitation(ctx context.Context, userID, code string) error
	CreateOrganization(ctx context.Context, name, ownerID string) (*roster.Organization, error)
	DeclineInvitation(ctx context.Context, code string) error
	InviteUserToOrganization(ctx context.Context, email, orgID string) (*roster.Invitation, error)
	RefreshInvitation(ctx context.Context, invitationID string) (*roster.Invitation, error)
}
