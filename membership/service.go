package membership

import (
	"context"
	"fmt"

	"github.com/haleyrc/cheevos/lib/db"
	"github.com/haleyrc/cheevos/lib/time"
)

type MembershipRepository interface {
	CreateMembership(ctx context.Context, tx db.Transaction, m *Membership) error
}

type MembershipService struct {
	DB   db.Database
	Repo MembershipRepository
}

func (ms *MembershipService) AddMemberToOrganization(ctx context.Context, userID, orgID string) error {
	mem := &Membership{
		OrganizationID: orgID,
		UserID:         userID,
		Joined:         time.Now(),
	}
	if err := mem.Validate(); err != nil {
		return fmt.Errorf("add member to organization failed: %w", err)
	}

	err := ms.DB.Call(ctx, func(ctx context.Context, tx db.Transaction) error {
		return ms.Repo.CreateMembership(ctx, tx, mem)
	})
	if err != nil {
		return fmt.Errorf("add member to organization failed: %w", err)
	}

	return nil
}
