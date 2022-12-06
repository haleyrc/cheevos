package domain

import (
	"context"

	"github.com/haleyrc/cheevos/internal/password"
)

type AuthService interface {
	GetUser(ctx context.Context, id string) (*User, error)
	SignUp(ctx context.Context, username string, password password.Password) (*User, error)
}
