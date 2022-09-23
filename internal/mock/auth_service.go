package mock

import (
	"context"

	"github.com/haleyrc/cheevos/auth"
)

type AuthService struct {
	SignUpFn func(ctx context.Context, username, password string) (*auth.User, error)
}

func (as *AuthService) SignUp(ctx context.Context, username, password string) (*auth.User, error) {
	return as.SignUpFn(ctx, username, password)
}
