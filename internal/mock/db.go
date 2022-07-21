package mock

import (
	"context"

	"github.com/haleyrc/cheevos"
)

var _ = cheevos.Database(NewDatabase())

// NewDatabase returns a mock database ready for use. This method is preferred
// to constructing the Database object manually.
func NewDatabase() *Database {
	return &Database{}
}

// Database is a mock implementation of both the
// [github.com/haleyrc/cheevos.Database] and
// [github.com/haleyrc/cheevos.Transaction] interfaces.
type Database struct {
	AddUserToOrganizationFn func(ctx context.Context, orgID, userID string) error
	AwardCheevoToUserFn     func(ctx context.Context, cheevoID, userID string) error
	GetCheevoFn             func(ctx context.Context, cheevoID string) (*cheevos.Cheevo, error)
	GetOrganizationFn       func(ctx context.Context, orgID string) (*cheevos.Organization, error)
	GetUserFn               func(ctx context.Context, userID string) (*cheevos.User, error)
}

// Call fully implements the [github.com/haleyrc/cheevos.Database] interface.
func (db *Database) Call(ctx context.Context, f func(ctx context.Context, tx cheevos.Transaction) error) error {
	return f(ctx, db)
}

// AddUserToOrganization partially implements the
// [github.com/haleyrc/cheevos.Transaction] interface.
func (db *Database) AddUserToOrganization(ctx context.Context, orgID, userID string) error {
	if db.AddUserToOrganizationFn == nil {
		panic("AddUserToOrganizationFn is not defined. Please define an implementation in your test to use the mock database.")
	}
	return db.AddUserToOrganizationFn(ctx, orgID, userID)
}

// AwardCheevoToUser partially implements the
// [github.com/haleyrc/cheevos.Transaction] interface.
func (db *Database) AwardCheevoToUser(ctx context.Context, cheevoID, userID string) error {
	if db.AwardCheevoToUserFn == nil {
		panic("AwardCheevoToUserFn is not defined. Please define an implementation in your test to use the mock database.")
	}
	return db.AwardCheevoToUserFn(ctx, cheevoID, userID)
}

// GetCheevo partially implements the [github.com/haleyrc/cheevos.Transaction]
// interface.
func (db *Database) GetCheevo(ctx context.Context, cheevoID string) (*cheevos.Cheevo, error) {
	if db.GetCheevoFn == nil {
		panic("GetCheevoFn is not defined. Please define an implementation in your test to use the mock database.")
	}
	return db.GetCheevoFn(ctx, cheevoID)
}

// GetOrganization partially implements the
// [github.com/haleyrc/cheevos.Transaction] interface.
func (db *Database) GetOrganization(ctx context.Context, orgID string) (*cheevos.Organization, error) {
	if db.GetOrganizationFn == nil {
		panic("GetOrganizationFn is not defined. Please define an implementation in your test to use the mock database.")
	}
	return db.GetOrganizationFn(ctx, orgID)
}

// GetUser partially implements the [github.com/haleyrc/cheevos.Transaction]
// interface.
func (db *Database) GetUser(ctx context.Context, userID string) (*cheevos.User, error) {
	if db.GetUserFn == nil {
		panic("GetUserFn is not defined. Please define an implementation in your test to use the mock database.")
	}
	return db.GetUserFn(ctx, userID)
}
