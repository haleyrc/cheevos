package mock

import (
	"context"

	"github.com/haleyrc/cheevos"
)

var _ cheevos.AuthService = &AuthService{}

type AuthService struct {
	GetUserFn func(ctx context.Context, id string) (*cheevos.User, error)
	SignUpFn  func(ctx context.Context, username, password string) (*cheevos.User, error)
}

func (as *AuthService) GetUser(ctx context.Context, id string) (*cheevos.User, error) {
	return as.GetUserFn(ctx, id)
}

func (as *AuthService) SignUp(ctx context.Context, username, password string) (*cheevos.User, error) {
	return as.SignUpFn(ctx, username, password)
}
