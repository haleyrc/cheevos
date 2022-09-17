package mock

import (
	"context"

	"github.com/haleyrc/cheevos/lib/db"
	"github.com/haleyrc/cheevos/user"
)

type CreateUserArgs struct {
	User         *user.User
	PasswordHash string
}

type UserRepository struct {
	CreateUserFn     func(ctx context.Context, tx db.Transaction, u *user.User, hashedPassword string) error
	CreateUserCalled struct {
		Count int
		With  CreateUserArgs
	}
}

func (ur *UserRepository) CreateUser(ctx context.Context, tx db.Transaction, u *user.User, hashedPassword string) error {
	if ur.CreateUserFn == nil {
		return mockMethodNotDefined("CreateUser")
	}
	ur.CreateUserCalled.Count++
	ur.CreateUserCalled.With = CreateUserArgs{User: u, PasswordHash: hashedPassword}
	return ur.CreateUserFn(ctx, tx, u, hashedPassword)
}
