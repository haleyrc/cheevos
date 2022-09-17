package invitation

import (
	"context"
	"fmt"

	"github.com/haleyrc/cheevos/lib/db"
	"github.com/haleyrc/cheevos/lib/hash"
	"github.com/haleyrc/cheevos/lib/random"
	"github.com/haleyrc/cheevos/lib/time"
)

const InvitationValidFor = time.Hour
const CodeLength = 32

type Emailer interface {
	SendInvitation(ctx context.Context, email, code string) error
}

type InvitationRepository interface {
	CreateInvitation(ctx context.Context, tx db.Transaction, i *Invitation, hashedCode string) error
}

type InvitationService struct {
	DB    db.Database
	Email Emailer
	Repo  InvitationRepository
}

func (is *InvitationService) InviteUserToOrganization(ctx context.Context, email, orgID string) (*Invitation, error) {
	invitation := &Invitation{
		Email:          email,
		OrganizationID: orgID,
		Expires:        time.Now().Add(InvitationValidFor),
	}
	if err := invitation.Validate(); err != nil {
		return nil, fmt.Errorf("invite user to organization failed: %w", err)
	}

	code := random.String(CodeLength)
	hashedCode := hash.Generate(code)

	err := is.DB.Call(ctx, func(ctx context.Context, tx db.Transaction) error {
		if err := is.Repo.CreateInvitation(ctx, tx, invitation, hashedCode); err != nil {
			return err
		}

		return is.Email.SendInvitation(ctx, invitation.Email, code)
	})
	if err != nil {
		return nil, fmt.Errorf("invite user to organization failed: %w", err)
	}

	return invitation, nil
}
