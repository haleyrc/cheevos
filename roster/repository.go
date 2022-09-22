package roster

import (
	"context"
	"fmt"

	"github.com/haleyrc/cheevos/lib/db"
)

type Repository struct{}

func (repo *Repository) CreateInvitation(ctx context.Context, tx db.Tx, i *Invitation, hashedCode string) error {
	return fmt.Errorf("TODO")
}

func (repo *Repository) CreateMembership(ctx context.Context, tx db.Tx, m *Membership) error {
	query := `INSERT INTO memberships (organization_id, user_id, joined_at) VALUES ($1, $2, $3);`
	if err := tx.Exec(ctx, query, m.OrganizationID, m.UserID, m.Joined); err != nil {
		return fmt.Errorf("create membership failed: %w", err)
	}
	return nil
}

func (repo *Repository) CreateOrganization(ctx context.Context, tx db.Tx, o *Organization) error {
	query := `INSERT INTO organizations (id, name, owner_id) VALUES ($1, $2, $3);`
	if err := tx.Exec(ctx, query, o.ID, o.Name, o.OwnerID); err != nil {
		return fmt.Errorf("create organization failed: %w", err)
	}
	return nil
}

func (repo *Repository) DeleteInvitationByCode(ctx context.Context, tx db.Tx, hashedCode string) error {
	return fmt.Errorf("TODO")
}

func (repo *Repository) GetInvitationByCode(ctx context.Context, tx db.Tx, hashedCode string) (*Invitation, error) {
	return nil, fmt.Errorf("TODO")
}

func (repo *Repository) SaveInvitation(ctx context.Context, tx db.Tx, i *Invitation, hashedCode string) error {
	return fmt.Errorf("TODO")
}
