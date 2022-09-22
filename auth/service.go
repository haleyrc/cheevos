package auth

import (
	"context"
	"fmt"

	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/lib/db"
	"github.com/haleyrc/cheevos/lib/hash"
)

// Service represents the main entrypoint for managing user.
type Service struct {
	DB   db.Database
	Repo interface {
		CreateUser(ctx context.Context, tx db.Tx, u *User, hashedPassword string) error
	}
}

// SignUp creates a new user and persists it to the database. It returns a
// response containing the new organization if successful.
func (us *Service) SignUp(ctx context.Context, username, password string) (*User, error) {
	var user *User
	err := us.DB.WithTx(ctx, func(ctx context.Context, tx db.Tx) error {
		user = &User{
			ID:       uuid.New(),
			Username: username,
		}
		if err := user.Validate(); err != nil {
			return err
		}

		hashedPassword := hash.Generate(password)

		return us.Repo.CreateUser(ctx, tx, user, hashedPassword)
	})
	if err != nil {
		return nil, fmt.Errorf("sign up failed: %w", err)
	}

	return user, nil
}
