package cheevos

import (
	"context"
)

type AuthService interface {
	GetUser(ctx context.Context, id string) (*User, error)
	SignUp(ctx context.Context, username, password string) (*User, error)
}
