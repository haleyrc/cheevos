package pg

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/haleyrc/cheevos"
)

func Port(port int) string {
	return fmt.Sprintf(":%d", port)
}

type Parameters struct {
	Database string
	Host     string
	Password string
	Port     string

	// SSL enables or disables SSL connections to the database.
	//
	// Note: This option is currently not used.
	SSL bool

	Username string
}

func (p Parameters) String() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s%s/%s?sslmode=disable",
		p.Username,
		p.Password,
		p.Host,
		p.Port,
		p.Database,
	)
}

func Connect(ctx context.Context, conn string) (*Database, error) {
	db, err := sqlx.ConnectContext(ctx, "postgres", conn)
	if err != nil {
		return nil, fmt.Errorf("connect failed: %w", err)
	}
	return &Database{db: db}, nil
}

func ConnectWithRetries(ctx context.Context, retries int, conn string) (*Database, error) {
	var db *sqlx.DB
	var err error
	var i int

	for {
		if i+1 > retries {
			return nil, fmt.Errorf("connect failed: %w", err)
		}

		if i != 0 {
			wait := time.Second * (1 << i)
			time.Sleep(wait)
		}

		db, err = sqlx.ConnectContext(ctx, "postgres", conn)
		if err == nil {
			return &Database{db: db}, nil
		}

		i++
	}
}

type Database struct {
	db *sqlx.DB
}

func (db *Database) Call(context.Context, func(ctx context.Context, tx Transaction) error) error {
	return fmt.Errorf("TODO")
}

func (db *Database) Close() error {
	if err := db.db.Close(); err != nil {
		return fmt.Errorf("close failed: %w", err)
	}
	return nil
}

func (db *Database) Ping() error {
	if err := db.db.Ping(); err != nil {
		return fmt.Errorf("ping failed: %w", err)
	}
	return nil
}

type Transaction struct {
}

func (tx *Transaction) AwardCheevoToUser(ctx context.Context, cheevoID, userID string) error {
	return fmt.Errorf("TODO")
}

func (tx *Transaction) GetCheevo(ctx context.Context, cheevoID string) (*cheevos.Cheevo, error) {
	return nil, fmt.Errorf("TODO")
}

func (tx *Transaction) GetOrganization(ctx context.Context, orgID string) (*cheevos.Organization, error) {
	return nil, fmt.Errorf("TODO")
}

func (tx *Transaction) GetUser(ctx context.Context, userID string) (*cheevos.User, error) {
	return nil, fmt.Errorf("TODO")
}

func (tx *Transaction) AddUserToOrganization(ctx context.Context, orgID, userID string) error {
	return fmt.Errorf("TODO")
}
