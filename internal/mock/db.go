package mock

import (
	"context"
	"fmt"

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
	AddUserToOrganizationFn func(ctx context.Context, org *cheevos.Organization, user *cheevos.User) error
	AwardCheevoToUserFn     func(ctx context.Context, cheevo *cheevos.Cheevo, awardee, awarder *cheevos.User) error
	CreateOrganizationFn    func(ctx context.Context, org *cheevos.Organization) error
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
func (db *Database) AddUserToOrganization(ctx context.Context, org *cheevos.Organization, user *cheevos.User) error {
	if db.AddUserToOrganizationFn == nil {
		panicNotDefined("AddUserToOrganizationFn")
	}
	return db.AddUserToOrganizationFn(ctx, org, user)
}

// AwardCheevoToUser partially implements the
// [github.com/haleyrc/cheevos.Transaction] interface.
func (db *Database) AwardCheevoToUser(ctx context.Context, cheevo *cheevos.Cheevo, awardee, awarder *cheevos.User) error {
	if db.AwardCheevoToUserFn == nil {
		panicNotDefined("AwardCheevoToUserFn")
	}
	return db.AwardCheevoToUserFn(ctx, cheevo, awardee, awarder)
}

func (db *Database) CreateCheevo(ctx context.Context, cheevo *cheevos.Cheevo) error {
	return fmt.Errorf("TODO")
}

func (db *Database) CreateOrganization(ctx context.Context, org *cheevos.Organization) error {
	if db.CreateOrganizationFn == nil {
		panicNotDefined("CreateOrganizationFn")
	}
	return db.CreateOrganizationFn(ctx, org)
}

func (db *Database) CreateUser(ctx context.Context, user *cheevos.User) error {
	return fmt.Errorf("TODO")
}

// GetCheevo partially implements the [github.com/haleyrc/cheevos.Transaction]
// interface.
func (db *Database) GetCheevo(ctx context.Context, cheevoID string) (*cheevos.Cheevo, error) {
	if db.GetCheevoFn == nil {
		panicNotDefined("GetCheevoFn")
	}
	return db.GetCheevoFn(ctx, cheevoID)
}

// GetOrganization partially implements the
// [github.com/haleyrc/cheevos.Transaction] interface.
func (db *Database) GetOrganization(ctx context.Context, orgID string) (*cheevos.Organization, error) {
	if db.GetOrganizationFn == nil {
		panicNotDefined("GetOrganizationFn")
	}
	return db.GetOrganizationFn(ctx, orgID)
}

// GetUser partially implements the [github.com/haleyrc/cheevos.Transaction]
// interface.
func (db *Database) GetUser(ctx context.Context, userID string) (*cheevos.User, error) {
	if db.GetUserFn == nil {
		panicNotDefined("GetUserFn")
	}
	return db.GetUserFn(ctx, userID)
}

func panicNotDefined(name string) {
	panic(fmt.Sprintf("%s is not defined. Please define an implementation in your test to use the mock database.", name))
}
