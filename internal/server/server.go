package server

import (
	"context"
)

type AuthorizationService interface {
	CanAwardCheevo(ctx context.Context, awarderID, recipientID, cheevoID string) error
	CanCreateCheevo(ctx context.Context, userID, orgID string) error
	CanGetCheevo(ctx context.Context, userID, cheevoID string) error
	CanInviteUsersToOrganization(ctx context.Context, userID, orgID string) error
	CanRefreshInvitation(ctx context.Context, userID, invitationID string) error
}
