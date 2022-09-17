package user

import (
	"context"
	"fmt"

	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/lib/db"
	"github.com/haleyrc/cheevos/lib/hash"
)

type UserRepository interface {
	CreateUser(ctx context.Context, tx db.Transaction, u *User, hashedPassword string) error
}

// UserService represents the main entrypoint for managing user.
type UserService struct {
	DB   db.Database
	Repo UserRepository
}

// SignUp creates a new user and persists it to the database. It returns a
// response containing the new organization if successful.
func (us *UserService) SignUp(ctx context.Context, username, password string) (*User, error) {
	user := &User{
		ID:       uuid.New(),
		Username: username,
	}
	if err := user.Validate(); err != nil {
		return nil, fmt.Errorf("sign up failed: %w", err)
	}

	err := us.DB.Call(ctx, func(ctx context.Context, tx db.Transaction) error {
		hashedPassword := hash.Generate(password)
		return us.Repo.CreateUser(ctx, tx, user, hashedPassword)
	})
	if err != nil {
		return nil, fmt.Errorf("sign up failed: %w", err)
	}

	return user, nil
}
