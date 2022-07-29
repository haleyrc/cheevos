package pg

import (
	"context"
	"fmt"

	"github.com/haleyrc/cheevos"
	"github.com/jmoiron/sqlx"
)

type Transaction struct {
	tx *sqlx.Tx
}

func (tx *Transaction) AddUserToOrganization(ctx context.Context, org *cheevos.Organization, user *cheevos.User) error {
	query := `INSERT INTO memberships (organization_id, user_id) VALUES ($1, $2);`
	_, err := tx.tx.ExecContext(ctx, query, org.ID, user.ID)
	if err != nil {
		return fmt.Errorf("add user to organization failed: %w", err)
	}
	return nil
}

func (tx *Transaction) AwardCheevoToUser(ctx context.Context, cheevo *cheevos.Cheevo, awardee, awarder *cheevos.User) error {
	query := `INSERT INTO awards (cheevo_id, awardee_id, awarder_id) VALUES ($1, $2, $3);`
	_, err := tx.tx.ExecContext(ctx, query, cheevo.ID, awardee.ID, awarder.ID)
	if err != nil {
		return fmt.Errorf("award cheevo to user failed: %w", err)
	}
	return nil
}

func (tx *Transaction) CreateCheevo(ctx context.Context, cheevo *cheevos.Cheevo) error {
	query := `INSERT INTO cheevos (id, name, description, organization_id) VALUES ($1, $2, $3, $4);`
	_, err := tx.tx.ExecContext(ctx, query, cheevo.ID, cheevo.Name, cheevo.Description, cheevo.Organization)
	if err != nil {
		return fmt.Errorf("create cheevo failed: %w", err)
	}
	return nil
}

func (tx *Transaction) CreateOrganization(ctx context.Context, org *cheevos.Organization) error {
	query := `INSERT INTO organizations (id, name, owner) VALUES ($1, $2, $3);`
	_, err := tx.tx.ExecContext(ctx, query, org.ID, org.Name, org.Owner)
	if err != nil {
		return fmt.Errorf("create organization failed: %w", err)
	}
	return nil
}

func (tx *Transaction) CreateUser(ctx context.Context, user *cheevos.User) error {
	query := `INSERT INTO users (id, username) VALUES ($1, $2);`
	_, err := tx.tx.ExecContext(ctx, query, user.ID, user.Username)
	if err != nil {
		return fmt.Errorf("create user failed: %w", err)
	}
	return nil
}

func (tx *Transaction) GetCheevo(ctx context.Context, cheevoID string) (*cheevos.Cheevo, error) {
	query := `SELECT id, name, description, organization_id FROM cheevos WHERE id = $1;`

	var cheevo cheevos.Cheevo
	if err := tx.tx.QueryRowxContext(ctx, query, cheevoID).Scan(&cheevo.ID, &cheevo.Name, &cheevo.Description, &cheevo.Organization); err != nil {
		return nil, fmt.Errorf("get cheevo failed: %w", err)
	}

	return &cheevo, nil
}

func (tx *Transaction) GetOrganization(ctx context.Context, orgID string) (*cheevos.Organization, error) {
	query := `SELECT id, name, owner FROM organizations WHERE id = $1;`

	var org cheevos.Organization
	if err := tx.tx.QueryRowxContext(ctx, query, orgID).Scan(&org.ID, &org.Name, &org.Owner); err != nil {
		return nil, fmt.Errorf("get organization failed: %w", err)
	}

	return &org, nil
}

func (tx *Transaction) GetUser(ctx context.Context, userID string) (*cheevos.User, error) {
	query := `SELECT id, username FROM users WHERE id = $1;`

	var user cheevos.User
	if err := tx.tx.QueryRowxContext(ctx, query, userID).Scan(&user.ID, &user.Username); err != nil {
		return nil, fmt.Errorf("get user failed: %w", err)
	}

	return &user, nil
}
