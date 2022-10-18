package repository

import (
	"context"
	"fmt"

	"github.com/haleyrc/cheevos"
	"github.com/haleyrc/cheevos/internal/lib/db"
	"github.com/haleyrc/cheevos/internal/repository/sql"
)

type RosterRepository struct{}

func (repo *RosterRepository) DeleteInvitationByCode(ctx context.Context, tx db.Tx, hashedCode string) error {
	if err := tx.Exec(ctx, sql.DeleteInvitationQuery, hashedCode); err != nil {
		return fmt.Errorf("delete invitation by code failed: %w", err)
	}
	return nil
}

func (repo *RosterRepository) GetInvitation(ctx context.Context, tx db.Tx, i *cheevos.Invitation, id string) error {
	if err := tx.QueryRow(ctx, sql.GetInvitationQuery, id).Scan(&i.Email, &i.OrganizationID, &i.Expires); err != nil {
		return fmt.Errorf("get invitation failed: %w", err)
	}
	return nil
}

func (repo *RosterRepository) GetInvitationByCode(ctx context.Context, tx db.Tx, i *cheevos.Invitation, hashedCode string) error {
	if err := tx.QueryRow(ctx, sql.GetInvitationByCodeQuery, hashedCode).Scan(&i.Email, &i.OrganizationID, &i.Expires); err != nil {
		return fmt.Errorf("get invitation by code failed: %w", err)
	}
	return nil
}

func (repo *RosterRepository) GetMembership(ctx context.Context, tx db.Tx, m *cheevos.Membership, orgID, userID string) error {
	if err := tx.QueryRow(ctx, sql.GetMembershipQuery, orgID, userID).Scan(&m.OrganizationID, &m.UserID, &m.Joined); err != nil {
		return fmt.Errorf("get member failed: %w", err)
	}
	return nil
}

func (repo *RosterRepository) InsertInvitation(ctx context.Context, tx db.Tx, i *cheevos.Invitation, hashedCode string) error {
	if err := tx.Exec(ctx, sql.InsertInvitationQuery, i.ID, i.Email, i.OrganizationID, i.Expires, hashedCode); err != nil {
		return fmt.Errorf("create invitation failed: %w", err)
	}
	return nil
}

func (repo *RosterRepository) InsertMembership(ctx context.Context, tx db.Tx, m *cheevos.Membership) error {
	if err := tx.Exec(ctx, sql.InsertMembershipQuery, m.OrganizationID, m.UserID, m.Joined); err != nil {
		return fmt.Errorf("create membership failed: %w", err)
	}
	return nil
}

func (repo *RosterRepository) InsertOrganization(ctx context.Context, tx db.Tx, o *cheevos.Organization) error {
	if err := tx.Exec(ctx, sql.InsertOrganizationQuery, o.ID, o.Name, o.OwnerID); err != nil {
		return fmt.Errorf("create organization failed: %w", err)
	}
	return nil
}

func (repo *RosterRepository) UpdateInvitation(ctx context.Context, tx db.Tx, i *cheevos.Invitation, hashedCode string) error {
	if err := tx.Exec(ctx, sql.UpdateInvitationQuery, i.Email, i.OrganizationID, i.Expires, hashedCode); err != nil {
		return fmt.Errorf("save invitation failed: %w", err)
	}
	return nil
}
