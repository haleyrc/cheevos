package user

import (
	"context"
	"fmt"

	"github.com/haleyrc/cheevos/lib/db"
)

type UserRepository interface {
	CreateUser(ctx context.Context, tx db.Transaction, u *User) error
}

type userRepository struct{}

func (ur *userRepository) CreateUser(ctx context.Context, tx db.Transaction, u *User) error {
	return fmt.Errorf("TODO")
}
