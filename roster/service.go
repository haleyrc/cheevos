package roster

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/core"
	"github.com/haleyrc/cheevos/lib/db"
	"github.com/haleyrc/cheevos/lib/hash"
	"github.com/haleyrc/cheevos/lib/random"
	"github.com/haleyrc/cheevos/lib/time"
)

var (
	InvitationValidFor = time.Hour
)

const (
	CodeLength = 32
)

type Emailer interface {
	SendInvitation(ctx context.Context, email, code string) error
}

type InvitationsRepository interface {
	DeleteInvitationByCode(ctx context.Context, tx db.Tx, hashedCode string) error
	GetInvitation(ctx context.Context, tx db.Tx, i *Invitation, id string) error
	GetInvitationByCode(ctx context.Context, tx db.Tx, i *Invitation, hashedCode string) error
	InsertInvitation(ctx context.Context, tx db.Tx, i *Invitation, hashedCode string) error
	UpdateInvitation(ctx context.Context, tx db.Tx, i *Invitation, hashedCode string) error
}

type MembershipsRepository interface {
	GetMembership(ctx context.Context, tx db.Tx, m *Membership, orgID, userID string) error
	InsertMembership(ctx context.Context, tx db.Tx, m *Membership) error
}

type OrganizationsRepository interface {
	InsertOrganization(ctx context.Context, tx db.Tx, org *Organization) error
}

type Service struct {
	DB    db.Database
	Email Emailer
	Repo  interface {
		InvitationsRepository
		MembershipsRepository
		OrganizationsRepository
	}
}

func (svc *Service) AcceptInvitation(ctx context.Context, userID, code string) error {
	err := svc.DB.WithTx(ctx, func(ctx context.Context, tx db.Tx) error {
		var invitation Invitation

		hashedCode := hash.Generate(code)
		if err := svc.Repo.GetInvitationByCode(ctx, tx, &invitation, hashedCode); err != nil {
			return core.WrapError(err)
		}

		if invitation.Expired() {
			return core.NewRawError(http.StatusGone, "Your invitation has expired. Please contact your organization administrator for a new invitation.")
		}

		membership := &Membership{
			OrganizationID: invitation.OrganizationID,
			UserID:         userID,
			Joined:         time.Now(),
		}
		if err := membership.Validate(); err != nil {
			return core.WrapError(err)
		}
		if err := svc.Repo.InsertMembership(ctx, tx, membership); err != nil {
			return core.WrapError(err)
		}

		return svc.Repo.DeleteInvitationByCode(ctx, tx, hashedCode)
	})
	if err != nil {
		return fmt.Errorf("accept invitation failed: %w", err)
	}

	return nil
}

// CreateOrganization creates a new organization and persists it to the
// database. It returns a response containing the new organization if
// successful.
func (svc *Service) CreateOrganization(ctx context.Context, name, ownerID string) (*Organization, error) {
	var org Organization

	err := svc.DB.WithTx(ctx, func(ctx context.Context, tx db.Tx) error {
		org = Organization{
			ID:      uuid.New(),
			Name:    name,
			OwnerID: ownerID,
		}
		if err := org.Validate(); err != nil {
			return core.WrapError(err)
		}

		membership := &Membership{
			OrganizationID: org.ID,
			UserID:         ownerID,
			Joined:         time.Now(),
		}
		if err := membership.Validate(); err != nil {
			return core.WrapError(err)
		}

		if err := svc.Repo.InsertOrganization(ctx, tx, &org); err != nil {
			return core.WrapError(err)
		}

		return svc.Repo.InsertMembership(ctx, tx, membership)
	})
	if err != nil {
		return nil, fmt.Errorf("create organization failed: %w", err)
	}

	return &org, nil
}

func (svc *Service) DeclineInvitation(ctx context.Context, code string) error {
	err := svc.DB.WithTx(ctx, func(ctx context.Context, tx db.Tx) error {
		return svc.Repo.DeleteInvitationByCode(ctx, tx, hash.Generate(code))
	})
	if err != nil {
		return fmt.Errorf("decline invitation failed: %w", err)
	}
	return nil
}

func (svc *Service) GetInvitation(ctx context.Context, id string) (*Invitation, error) {
	var invitation Invitation

	err := svc.DB.WithTx(ctx, func(ctx context.Context, tx db.Tx) error {
		return svc.Repo.GetInvitation(ctx, tx, &invitation, id)
	})
	if err != nil {
		return nil, fmt.Errorf("get invitation failed: %w", err)
	}

	return &invitation, nil
}

func (svc *Service) InviteUserToOrganization(ctx context.Context, email, orgID string) (*Invitation, error) {
	var invitation Invitation

	err := svc.DB.WithTx(ctx, func(ctx context.Context, tx db.Tx) error {
		invitation = Invitation{
			ID:             uuid.New(),
			Email:          email,
			OrganizationID: orgID,
			Expires:        time.Now().Add(InvitationValidFor),
		}
		if err := invitation.Validate(); err != nil {
			return core.WrapError(err)
		}

		code := random.String(CodeLength)
		if err := svc.Repo.InsertInvitation(ctx, tx, &invitation, hash.Generate(code)); err != nil {
			return core.WrapError(err)
		}

		return svc.Email.SendInvitation(ctx, invitation.Email, code)
	})
	if err != nil {
		return nil, fmt.Errorf("invite user to organization failed: %w", err)
	}

	return &invitation, nil
}

func (svc *Service) IsMember(ctx context.Context, orgID, userID string) error {
	err := svc.DB.WithTx(ctx, func(ctx context.Context, tx db.Tx) error {
		return svc.Repo.GetMembership(ctx, tx, &Membership{}, orgID, userID)
	})
	if err != nil {
		return fmt.Errorf("is member failed: %w", err)
	}
	return nil
}

// Refreshing an invitation will invalidate the initial invitation email (as
// well as any other refresh emails).
func (svc *Service) RefreshInvitation(ctx context.Context, id string) (*Invitation, error) {
	var invitation Invitation

	err := svc.DB.WithTx(ctx, func(ctx context.Context, tx db.Tx) error {
		if err := svc.Repo.GetInvitation(ctx, tx, &invitation, id); err != nil {
			return core.WrapError(err)
		}

		invitation.Expires = time.Now().Add(InvitationValidFor)

		newCode := random.String(CodeLength)
		if err := svc.Repo.UpdateInvitation(ctx, tx, &invitation, hash.Generate(newCode)); err != nil {
			return core.WrapError(err)
		}

		return svc.Email.SendInvitation(ctx, invitation.Email, newCode)
	})
	if err != nil {
		return nil, fmt.Errorf("refresh invitation failed: %w", err)
	}

	return &invitation, nil
}
