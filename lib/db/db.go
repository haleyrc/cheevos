package db

import "context"

// Database represents a connection to a persistence layer. A Database is only
// responsible for opening a transaction which is then responsible for all of
// the heavy lifting.
type Database interface {
	// WithTx is responsible for opening a transaction boundary on the parent
	// database and calling the provided function, passing the transaction in for
	// use by the service layer. If an error is returned from the function
	// closure, the transaction is rolled back and WithTx returns an error as
	// well.
	WithTx(context.Context, func(ctx context.Context, tx Tx) error) error
}

type Row interface {
	Scan(args ...interface{}) error
}

type Tx interface {
	Exec(ctx context.Context, query string, args ...interface{}) error
	QueryRow(ctx context.Context, query string, args ...interface{}) Row
}
