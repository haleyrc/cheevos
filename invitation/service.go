package invitation

import (
	"context"
	"fmt"

	"github.com/haleyrc/cheevos/lib/db"
	"github.com/haleyrc/cheevos/lib/hash"
	"github.com/haleyrc/cheevos/lib/random"
	"github.com/haleyrc/cheevos/lib/time"
)

var InvitationValidFor = time.Hour

const CodeLength = 32

type Emailer interface {
	SendInvitation(ctx context.Context, email, code string) error
}

type Repository interface {
	AddMemberToOrganization(ctx context.Context, tx db.Transaction, userID, orgID string) error
	CreateInvitation(ctx context.Context, tx db.Transaction, i *Invitation, hashedCode string) error
	DeleteInvitationByCode(ctx context.Context, tx db.Transaction, hashedCode string) error
	GetInvitationByCode(ctx context.Context, tx db.Transaction, hashedCode string) (*Invitation, error)
	SaveInvitation(ctx context.Context, tx db.Transaction, i *Invitation, hashedCode string) error
}

type Service struct {
	DB    db.Database
	Email Emailer
	Repo  Repository
}

func (is *Service) AcceptInvitation(ctx context.Context, userID, code string) error {
	err := is.DB.Call(ctx, func(ctx context.Context, tx db.Transaction) error {
		hashedCode := hash.Generate(code)
		invitation, err := is.Repo.GetInvitationByCode(ctx, tx, hashedCode)
		if err != nil {
			return err
		}

		if invitation.Expired() {
			return fmt.Errorf("invitation is expired")
		}

		if err := is.Repo.AddMemberToOrganization(ctx, tx, userID, invitation.OrganizationID); err != nil {
			return err
		}

		return is.Repo.DeleteInvitationByCode(ctx, tx, hashedCode)
	})
	if err != nil {
		return fmt.Errorf("accept invitation failed: %w", err)
	}

	return nil
}

func (is *Service) DeclineInvitation(ctx context.Context, code string) error {
	err := is.DB.Call(ctx, func(ctx context.Context, tx db.Transaction) error {
		hashedCode := hash.Generate(code)
		return is.Repo.DeleteInvitationByCode(ctx, tx, hashedCode)
	})
	if err != nil {
		return fmt.Errorf("decline invitation failed: %w", err)
	}

	return nil
}

func (is *Service) InviteUserToOrganization(ctx context.Context, email, orgID string) (*Invitation, error) {
	var invitation *Invitation
	err := is.DB.Call(ctx, func(ctx context.Context, tx db.Transaction) error {
		code := random.String(CodeLength)
		hashedCode := hash.Generate(code)

		invitation = &Invitation{
			Email:          email,
			OrganizationID: orgID,
			Expires:        time.Now().Add(InvitationValidFor),
		}
		if err := invitation.Validate(); err != nil {
			return err
		}

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

// Refreshing an invitation will invalidate the initial invitation email (as
// well as any other refresh emails).
func (is *Service) RefreshInvitation(ctx context.Context, code string) error {
	err := is.DB.Call(ctx, func(ctx context.Context, tx db.Transaction) error {
		hashedCode := hash.Generate(code)

		invitation, err := is.Repo.GetInvitationByCode(ctx, tx, hashedCode)
		if err != nil {
			return err
		}

		newCode := random.String(CodeLength)
		newCodeHash := hash.Generate(newCode)
		invitation.Expires = time.Now().Add(InvitationValidFor)

		if err := is.Repo.SaveInvitation(ctx, tx, invitation, newCodeHash); err != nil {
			return err
		}

		return is.Email.SendInvitation(ctx, invitation.Email, newCode)
	})
	if err != nil {
		return fmt.Errorf("refresh invitation failed: %w", err)
	}

	return nil
}
