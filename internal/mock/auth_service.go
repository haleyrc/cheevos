package mock

import (
	"context"

	"github.com/haleyrc/cheevos/domain"
	"github.com/haleyrc/cheevos/internal/password"
)

var _ domain.AuthService = &AuthService{}

type AuthService struct {
	GetUserFn func(ctx context.Context, id string) (*domain.User, error)
	SignUpFn  func(ctx context.Context, username string, password password.Password) (*domain.User, error)
}

func (as *AuthService) GetUser(ctx context.Context, id string) (*domain.User, error) {
	return as.GetUserFn(ctx, id)
}

func (as *AuthService) SignUp(ctx context.Context, username string, password password.Password) (*domain.User, error) {
	return as.SignUpFn(ctx, username, password)
}
