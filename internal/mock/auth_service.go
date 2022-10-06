package mock

import (
	"context"

	"github.com/haleyrc/cheevos/internal/server"
	"github.com/haleyrc/cheevos/internal/service"
)

var _ server.AuthenticationService = &AuthService{}

type AuthService struct {
	SignUpFn func(ctx context.Context, username, password string) (*service.User, error)
}

func (as *AuthService) SignUp(ctx context.Context, username, password string) (*service.User, error) {
	return as.SignUpFn(ctx, username, password)
}
