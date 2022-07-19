package cheevos

import (
	"context"
	"fmt"
	"strings"
)

type UserService struct{}

type SignUpRequest struct {
	Username string
	Password string
}

type SignUpResponse struct{}

func (us *UserService) SignUp(ctx context.Context, req SignUpRequest) (*SignUpResponse, error) {
	return nil, nil
}

type User struct{}
