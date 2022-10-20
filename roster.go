package cheevos

import (
	"context"

	"github.com/haleyrc/pkg/time"
)

var InvitationValidFor = time.Hour

const InvitationCodeLength = 32

type RosterService interface {
	AcceptInvitation(ctx context.Context, userID, code string) error
	CreateOrganization(ctx context.Context, name, ownerID string) (*Organization, error)
	DeclineInvitation(ctx context.Context, code string) error
	GetInvitation(ctx context.Context, id string) (*Invitation, error)
	InviteUserToOrganization(ctx context.Context, email, orgID string) (*Invitation, error)
	IsMember(ctx context.Context, orgID, userID string) error
	RefreshInvitation(ctx context.Context, id string) (*Invitation, error)
}
