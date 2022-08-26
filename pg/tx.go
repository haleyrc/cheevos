package pg

import (
	"context"
	"fmt"

	"github.com/haleyrc/cheevos"
	"github.com/haleyrc/cheevos/log"
	"github.com/jmoiron/sqlx"
)

type Transaction struct {
	logger log.Logger
	tx     *sqlx.Tx
}

func (tx *Transaction) AddUserToOrganization(ctx context.Context, org *cheevos.Organization, user *cheevos.User) error {
	tx.logger.Debug(ctx, "adding user to organization", log.Fields{"Organization": org, "User": user})
	query := `INSERT INTO memberships (organization_id, user_id) VALUES ($1, $2);`
	_, err := tx.tx.ExecContext(ctx, query, org.ID, user.ID)
	if err != nil {
		return fmt.Errorf("add user to organization failed: %w", err)
	}
	return nil
}

func (tx *Transaction) AwardCheevoToUser(ctx context.Context, cheevo *cheevos.Cheevo, awardee, awarder *cheevos.User) error {
	tx.logger.Debug(ctx, "awarding cheevo to user", log.Fields{"Cheevo": cheevo, "Awardee": awardee, "Awarder": awarder})
	query := `INSERT INTO awards (cheevo_id, awardee_id, awarder_id) VALUES ($1, $2, $3);`
	_, err := tx.tx.ExecContext(ctx, query, cheevo.ID, awardee.ID, awarder.ID)
	if err != nil {
		return fmt.Errorf("award cheevo to user failed: %w", err)
	}
	return nil
}

func (tx *Transaction) CreateCheevo(ctx context.Context, cheevo *cheevos.Cheevo) error {
	tx.logger.Debug(ctx, "creating cheevo", log.Fields{"Cheevo": cheevo})
	query := `INSERT INTO cheevos (id, name, description, organization_id) VALUES ($1, $2, $3, $4);`
	_, err := tx.tx.ExecContext(ctx, query, cheevo.ID, cheevo.Name, cheevo.Description, cheevo.Organization)
	if err != nil {
		return fmt.Errorf("create cheevo failed: %w", err)
	}
	return nil
}

func (tx *Transaction) CreateOrganization(ctx context.Context, org *cheevos.Organization) error {
	tx.logger.Debug(ctx, "creating organization", log.Fields{"Organization": org})
	query := `INSERT INTO organizations (id, name, owner) VALUES ($1, $2, $3);`
	_, err := tx.tx.ExecContext(ctx, query, org.ID, org.Name, org.Owner)
	if err != nil {
		return fmt.Errorf("create organization failed: %w", err)
	}
	return nil
}

func (tx *Transaction) CreateUser(ctx context.Context, user *cheevos.User) error {
	tx.logger.Debug(ctx, "creating user", log.Fields{"User": user})
	query := `INSERT INTO users (id, username) VALUES ($1, $2);`
	_, err := tx.tx.ExecContext(ctx, query, user.ID, user.Username)
	if err != nil {
		return fmt.Errorf("create user failed: %w", err)
	}
	return nil
}

func (tx *Transaction) GetCheevo(ctx context.Context, id string) (*cheevos.Cheevo, error) {
	tx.logger.Debug(ctx, "getting cheevo", log.Fields{"ID": id})
	query := `SELECT id, name, description, organization_id FROM cheevos WHERE id = $1;`

	var cheevo cheevos.Cheevo
	if err := tx.tx.QueryRowxContext(ctx, query, id).Scan(&cheevo.ID, &cheevo.Name, &cheevo.Description, &cheevo.Organization); err != nil {
		return nil, fmt.Errorf("get cheevo failed: %w", err)
	}

	return &cheevo, nil
}

func (tx *Transaction) GetOrganization(ctx context.Context, id string) (*cheevos.Organization, error) {
	tx.logger.Debug(ctx, "getting organization", log.Fields{"ID": id})
	query := `SELECT id, name, owner FROM organizations WHERE id = $1;`

	var org cheevos.Organization
	if err := tx.tx.QueryRowxContext(ctx, query, id).Scan(&org.ID, &org.Name, &org.Owner); err != nil {
		return nil, fmt.Errorf("get organization failed: %w", err)
	}

	return &org, nil
}

func (tx *Transaction) GetUser(ctx context.Context, id string) (*cheevos.User, error) {
	tx.logger.Debug(ctx, "getting user", log.Fields{"ID": id})
	query := `SELECT id, username FROM users WHERE id = $1;`

	var user cheevos.User
	if err := tx.tx.QueryRowxContext(ctx, query, id).Scan(&user.ID, &user.Username); err != nil {
		return nil, fmt.Errorf("get user failed: %w", err)
	}

	return &user, nil
}
