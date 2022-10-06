package server

import (
	"context"

	"github.com/haleyrc/cheevos/internal/service"
)

type AuthorizationService interface {
	CanAwardCheevo(ctx context.Context, awarderID, recipientID, cheevoID string) error
	CanCreateCheevo(ctx context.Context, userID, orgID string) error
	CanGetCheevo(ctx context.Context, userID, cheevoID string) error
	CanInviteUsersToOrganization(ctx context.Context, userID, orgID string) error
	CanRefreshInvitation(ctx context.Context, userID, invitationID string) error
}

type AuthenticationService interface {
	SignUp(ctx context.Context, username, password string) (*service.User, error)
}

type CheevosService interface {
	AwardCheevoToUser(ctx context.Context, recipientID, cheevoID string) error
	CreateCheevo(ctx context.Context, name, description, orgID string) (*service.Cheevo, error)
	GetCheevo(ctx context.Context, id string) (*service.Cheevo, error)
}

type RosterService interface {
	AcceptInvitation(ctx context.Context, userID, code string) error
	CreateOrganization(ctx context.Context, name, ownerID string) (*service.Organization, error)
	DeclineInvitation(ctx context.Context, code string) error
	InviteUserToOrganization(ctx context.Context, email, orgID string) (*service.Invitation, error)
	RefreshInvitation(ctx context.Context, invitationID string) (*service.Invitation, error)
}
