package cheevos

import "context"

// Database represents a connection to a persistence layer. A Database is only
// responsible for opening a transaction which is then responsible for all of
// the heavy lifting.
type Database interface {
	// Call is responsible for opening a transaction boundary on the parent
	// database and calling the provided function, passing the transaction in for
	// use by the service layer. If an error is returned from the function
	// closure, the transaction is rolled back and Call returns an error as well.
	Call(context.Context, func(ctx context.Context, tx Transaction) error) error
}

// Transaction is our primary handle for running persistence methods.
type Transaction interface {
	GetOrganization(ctx context.Context, orgID string) (*Organization, error)
	GetUser(ctx context.Context, userID string) (*User, error)
	AddUserToOrganization(ctx context.Context, orgID, userID string) error
}
