package mock

import (
	"context"

	"github.com/haleyrc/cheevos/user"
)

type UserService struct {
	SignUpFn func(ctx context.Context, username, password string) (*user.User, error)
}

func (us *UserService) SignUp(ctx context.Context, username, password string) (*user.User, error) {
	return us.SignUpFn(ctx, username, password)
}
