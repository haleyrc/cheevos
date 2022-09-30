package auth

import (
	"context"
	"fmt"

	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/core"
	"github.com/haleyrc/cheevos/lib/db"
	"github.com/haleyrc/cheevos/lib/hash"
)

type UsersRepository interface {
	CreateUser(ctx context.Context, tx db.Tx, u *User, hashedPassword string) error
	GetUser(ctx context.Context, tx db.Tx, u *User, id string) error
}

type Service struct {
	DB   db.Database
	Repo interface {
		UsersRepository
	}
}

func (svc *Service) GetUser(ctx context.Context, id string) (*User, error) {
	var user User

	err := svc.DB.WithTx(ctx, func(ctx context.Context, tx db.Tx) error {
		return svc.Repo.GetUser(ctx, tx, &user, id)
	})
	if err != nil {
		return nil, fmt.Errorf("get user failed: %w", err)
	}

	return &user, nil
}

// SignUp creates a new user and persists it to the database. It returns a
// response containing the new organization if successful.
func (svc *Service) SignUp(ctx context.Context, username, password string) (*User, error) {
	var user User

	err := svc.DB.WithTx(ctx, func(ctx context.Context, tx db.Tx) error {
		user = User{
			ID:       uuid.New(),
			Username: username,
		}
		if err := user.Validate(); err != nil {
			return core.WrapError(err)
		}
		return svc.Repo.CreateUser(ctx, tx, &user, hash.Generate(password))
	})
	if err != nil {
		return nil, fmt.Errorf("sign up failed: %w", err)
	}

	return &user, nil
}
