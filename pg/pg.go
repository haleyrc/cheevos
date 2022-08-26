package pg

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/haleyrc/cheevos"
	"github.com/haleyrc/cheevos/log"
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

type Options struct {
	ErrorFunc func(ctx context.Context, err error)
	Logger    log.Logger
}

func Connect(ctx context.Context, conn string, opts Options) (*Database, error) {
	sqlxDB, err := sqlx.ConnectContext(ctx, "postgres", conn)
	if err != nil {
		return nil, fmt.Errorf("connect failed: %w", err)
	}
	db := &Database{
		db:        sqlxDB,
		ErrorFunc: opts.ErrorFunc,
		Logger:    opts.Logger,
	}
	if db.ErrorFunc == nil {
		db.ErrorFunc = func(ctx context.Context, err error) {
			db.Logger.Error(ctx, "unexpected error", err)
		}
	}
	if db.Logger == nil {
		db.Logger = log.NullLogger{}
	}
	return db, nil
}

func ConnectWithRetries(ctx context.Context, retries int, conn string, opts Options) (*Database, error) {
	var i int
	var lastErr error
	for {
		if i+1 > retries {
			return nil, fmt.Errorf("connect failed: %w", lastErr)
		}

		if i != 0 {
			wait := time.Second * (1 << i)
			time.Sleep(wait)
		}

		db, err := Connect(ctx, conn, opts)
		if err == nil {
			return db, nil
		}
		lastErr = err

		i++
	}
}

type Database struct {
	ErrorFunc func(context.Context, error)
	Logger    log.Logger
	db        *sqlx.DB
}

func (db *Database) Call(ctx context.Context, f func(ctx context.Context, tx cheevos.Transaction) error) error {
	db.Logger.Debug(ctx, "begining transaction", nil)
	defer db.Logger.Debug(ctx, "ending transaction", nil)

	sqlxTx, err := db.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin transaction failed: %w", err)
	}

	if err := f(ctx, &Transaction{logger: db.Logger, tx: sqlxTx}); err != nil {
		rollbackErr := sqlxTx.Rollback()
		if rollbackErr != nil {
			db.ErrorFunc(ctx, rollbackErr)
		}
		return fmt.Errorf("transaction failed: %w", err)
	}

	if err := sqlxTx.Commit(); err != nil {
		return fmt.Errorf("commit failed: %w", err)
	}

	return nil
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
