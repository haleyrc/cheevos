package mock

import (
	"context"

	"github.com/haleyrc/cheevos"
)

type UserService struct {
	SignUpFn func(context.Context, cheevos.SignUpRequest) (*cheevos.SignUpResponse, error)
}

func (cs *UserService) SignUp(ctx context.Context, req cheevos.SignUpRequest) (*cheevos.SignUpResponse, error) {
	return cs.SignUpFn(ctx, req)
}
