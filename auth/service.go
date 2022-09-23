package auth

import (
	"context"
	"fmt"

	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/lib/db"
	"github.com/haleyrc/cheevos/lib/hash"
)

type IRepository interface {
	CreateUser(ctx context.Context, tx db.Tx, u *User, hashedPassword string) error
	GetUser(ctx context.Context, tx db.Tx, u *User, id string) error
}

// Service represents the main entrypoint for managing user.
type Service struct {
	DB   db.Database
	Repo IRepository
}

func (us *Service) GetUser(ctx context.Context, id string) (*User, error) {
	var user User

	err := us.DB.WithTx(ctx, func(ctx context.Context, tx db.Tx) error {
		return us.Repo.GetUser(ctx, tx, &user, id)
	})
	if err != nil {
		return nil, fmt.Errorf("get user failed: %w", err)
	}

	return &user, nil
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
