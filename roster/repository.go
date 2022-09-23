package roster

import (
	"context"
	"fmt"

	"github.com/haleyrc/cheevos/lib/db"
)

type Repository struct{}

func (repo *Repository) CreateInvitation(ctx context.Context, tx db.Tx, i *Invitation, hashedCode string) error {
	query := `INSERT INTO invitations (email, organization_id, expires_at, hashed_code) VALUES ($1, $2, $3, $4);`
	if err := tx.Exec(ctx, query, i.Email, i.OrganizationID, i.Expires, hashedCode); err != nil {
		return fmt.Errorf("create invitation failed: %w", err)
	}
	return nil
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
	query := `DELETE FROM invitations WHERE hashed_code = $1;`
	if err := tx.Exec(ctx, query, hashedCode); err != nil {
		return fmt.Errorf("delete invitation by code failed: %w", err)
	}
	return nil
}

func (repo *Repository) GetInvitationByCode(ctx context.Context, tx db.Tx, hashedCode string) (*Invitation, error) {
	var i Invitation

	query := `SELECT email, organization_id, expires_at FROM invitations WHERE hashed_code = $1;`
	if err := tx.QueryRow(ctx, query, hashedCode).Scan(&i.Email, &i.OrganizationID, &i.Expires); err != nil {
		return nil, fmt.Errorf("get invitation by code failed: %w", err)
	}

	return &i, nil
}

func (repo *Repository) GetMember(ctx context.Context, tx db.Tx, m *Membership, orgID, userID string) error {
	query := `SELECT organization_id, user_id, joined_at FROM memberships WHERE organization_id = $1 AND user_id = $2;`
	if err := tx.QueryRow(ctx, query, orgID, userID).Scan(&m.OrganizationID, &m.UserID, &m.Joined); err != nil {
		return fmt.Errorf("get member failed: %w", err)
	}
	return nil
}

func (repo *Repository) SaveInvitation(ctx context.Context, tx db.Tx, i *Invitation, hashedCode string) error {
	query := `UPDATE invitations SET expires_at = $3, hashed_code = $4 WHERE email = $1 AND organization_id = $2;`
	if err := tx.Exec(ctx, query, i.Email, i.OrganizationID, i.Expires, hashedCode); err != nil {
		return fmt.Errorf("save invitation failed: %w", err)
	}
	return nil
}
