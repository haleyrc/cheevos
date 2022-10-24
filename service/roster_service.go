package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/haleyrc/pkg/errors"
	"github.com/haleyrc/pkg/hash"
	"github.com/haleyrc/pkg/logger"
	"github.com/haleyrc/pkg/pg"
	"github.com/haleyrc/pkg/random"
	"github.com/haleyrc/pkg/time"
	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/domain"
)

var _ domain.RosterService = &rosterService{}

type Emailer interface {
	SendInvitation(ctx context.Context, email, code string) error
}

type RosterRepository interface {
	DeleteInvitationByCode(ctx context.Context, tx pg.Tx, hashedCode string) error
	GetInvitation(ctx context.Context, tx pg.Tx, i *domain.Invitation, id string) error
	GetInvitationByCode(ctx context.Context, tx pg.Tx, i *domain.Invitation, hashedCode string) error
	InsertInvitation(ctx context.Context, tx pg.Tx, i *domain.Invitation, hashedCode string) error
	UpdateInvitation(ctx context.Context, tx pg.Tx, i *domain.Invitation, hashedCode string) error
	GetMembership(ctx context.Context, tx pg.Tx, m *domain.Membership, orgID, userID string) error
	InsertMembership(ctx context.Context, tx pg.Tx, m *domain.Membership) error
	InsertOrganization(ctx context.Context, tx pg.Tx, org *domain.Organization) error
}

func NewRosterService(db Database, email Emailer, logger logger.Logger, repo RosterRepository) domain.RosterService {
	return &rosterLogger{
		Logger: logger,
		Service: &rosterService{
			DB:    db,
			Email: email,
			Repo:  repo,
		},
	}
}

type rosterService struct {
	DB    Database
	Email Emailer
	Repo  RosterRepository
}

func (svc *rosterService) AcceptInvitation(ctx context.Context, userID, code string) error {
	err := svc.DB.WithTx(ctx, func(ctx context.Context, tx pg.Tx) error {
		var invitation domain.Invitation

		hashedCode := hash.Generate(code)
		if err := svc.Repo.GetInvitationByCode(ctx, tx, &invitation, hashedCode); err != nil {
			return errors.WrapError(err)
		}

		if invitation.Expired() {
			return errors.NewRawError(http.StatusGone, "Your invitation has expired. Please contact your organization administrator for a new invitation.")
		}

		membership := &domain.Membership{
			OrganizationID: invitation.OrganizationID,
			UserID:         userID,
			Joined:         time.Now(),
		}
		if err := membership.Validate(); err != nil {
			return errors.WrapError(err)
		}
		if err := svc.Repo.InsertMembership(ctx, tx, membership); err != nil {
			return errors.WrapError(err)
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
func (svc *rosterService) CreateOrganization(ctx context.Context, name, ownerID string) (*domain.Organization, error) {
	var org domain.Organization

	err := svc.DB.WithTx(ctx, func(ctx context.Context, tx pg.Tx) error {
		org = domain.Organization{
			ID:      uuid.New(),
			Name:    name,
			OwnerID: ownerID,
		}
		if err := org.Validate(); err != nil {
			return errors.WrapError(err)
		}

		membership := &domain.Membership{
			OrganizationID: org.ID,
			UserID:         ownerID,
			Joined:         time.Now(),
		}
		if err := membership.Validate(); err != nil {
			return errors.WrapError(err)
		}

		if err := svc.Repo.InsertOrganization(ctx, tx, &org); err != nil {
			return errors.WrapError(err)
		}

		return svc.Repo.InsertMembership(ctx, tx, membership)
	})
	if err != nil {
		return nil, fmt.Errorf("create organization failed: %w", err)
	}

	return &org, nil
}

func (svc *rosterService) DeclineInvitation(ctx context.Context, code string) error {
	err := svc.DB.WithTx(ctx, func(ctx context.Context, tx pg.Tx) error {
		return svc.Repo.DeleteInvitationByCode(ctx, tx, hash.Generate(code))
	})
	if err != nil {
		return fmt.Errorf("decline invitation failed: %w", err)
	}
	return nil
}

func (svc *rosterService) GetInvitation(ctx context.Context, id string) (*domain.Invitation, error) {
	var invitation domain.Invitation

	err := svc.DB.WithTx(ctx, func(ctx context.Context, tx pg.Tx) error {
		return svc.Repo.GetInvitation(ctx, tx, &invitation, id)
	})
	if err != nil {
		return nil, fmt.Errorf("get invitation failed: %w", err)
	}

	return &invitation, nil
}

func (svc *rosterService) InviteUserToOrganization(ctx context.Context, email, orgID string) (*domain.Invitation, error) {
	var invitation domain.Invitation

	err := svc.DB.WithTx(ctx, func(ctx context.Context, tx pg.Tx) error {
		invitation = domain.Invitation{
			ID:             uuid.New(),
			Email:          email,
			OrganizationID: orgID,
			Expires:        time.Now().Add(domain.InvitationValidFor),
		}
		if err := invitation.Validate(); err != nil {
			return errors.WrapError(err)
		}

		code := random.String(domain.InvitationCodeLength)
		if err := svc.Repo.InsertInvitation(ctx, tx, &invitation, hash.Generate(code)); err != nil {
			return errors.WrapError(err)
		}

		return svc.Email.SendInvitation(ctx, invitation.Email, code)
	})
	if err != nil {
		return nil, fmt.Errorf("invite user to organization failed: %w", err)
	}

	return &invitation, nil
}

func (svc *rosterService) IsMember(ctx context.Context, orgID, userID string) error {
	err := svc.DB.WithTx(ctx, func(ctx context.Context, tx pg.Tx) error {
		return svc.Repo.GetMembership(ctx, tx, &domain.Membership{}, orgID, userID)
	})
	if err != nil {
		return fmt.Errorf("is member failed: %w", err)
	}
	return nil
}

// RefreshInvitation updates an invitation with a new expiration time and code.
// Refreshing an invitation will invalidate the initial invitation email (as
// well as any other refresh emails).
func (svc *rosterService) RefreshInvitation(ctx context.Context, id string) (*domain.Invitation, error) {
	var invitation domain.Invitation

	err := svc.DB.WithTx(ctx, func(ctx context.Context, tx pg.Tx) error {
		if err := svc.Repo.GetInvitation(ctx, tx, &invitation, id); err != nil {
			return errors.WrapError(err)
		}

		invitation.Expires = time.Now().Add(domain.InvitationValidFor)

		newCode := random.String(domain.InvitationCodeLength)
		if err := svc.Repo.UpdateInvitation(ctx, tx, &invitation, hash.Generate(newCode)); err != nil {
			return errors.WrapError(err)
		}

		return svc.Email.SendInvitation(ctx, invitation.Email, newCode)
	})
	if err != nil {
		return nil, fmt.Errorf("refresh invitation failed: %w", err)
	}

	return &invitation, nil
}
