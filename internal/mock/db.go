package mock

import (
	"context"

	"github.com/haleyrc/cheevos"
)

var _ = cheevos.Database(NewDatabase())

func NewDatabase() *Database {
	return &Database{}
}

type Database struct {
	AddUserToOrganizationFn func(ctx context.Context, orgID, userID string) error
	GetOrganizationFn       func(ctx context.Context, orgID string) (*cheevos.Organization, error)
	GetUserFn               func(ctx context.Context, userID string) (*cheevos.User, error)
}

func (db *Database) Call(ctx context.Context, f func(ctx context.Context, tx cheevos.Transaction) error) error {
	return f(ctx, db)
}

func (db *Database) AddUserToOrganization(ctx context.Context, orgID, userID string) error {
	if db.AddUserToOrganizationFn == nil {
		panic("AddUserToOrganizationFn is not defined. Please define an implementation in your test to use the mock database.")
	}
	return db.AddUserToOrganizationFn(ctx, orgID, userID)
}

func (db *Database) GetOrganization(ctx context.Context, orgID string) (*cheevos.Organization, error) {
	if db.GetOrganizationFn == nil {
		panic("GetOrganizationFn is not defined. Please define an implementation in your test to use the mock database.")
	}
	return db.GetOrganizationFn(ctx, orgID)
}

func (db *Database) GetUser(ctx context.Context, userID string) (*cheevos.User, error) {
	if db.GetUserFn == nil {
		panic("GetUserFn is not defined. Please define an implementation in your test to use the mock database.")
	}
	return db.GetUserFn(ctx, userID)
}
