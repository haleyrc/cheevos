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

func (req *SignUpRequest) normalize() {
	req.Username = strings.TrimSpace(req.Username)
	req.Password = strings.TrimSpace(req.Password)
}

type SignUpResponse struct{}

func (us *UserService) SignUp(ctx context.Context, req SignUpRequest) (*SignUpResponse, error) {
	req.normalize()

	if req.Username == "" {
		return nil, fmt.Errorf("invalid: username is blank")
	}

	if req.Password == "" {
		return nil, fmt.Errorf("invalid: password is blank")
	}

	return nil, nil
}

type User struct{}
