package cheevos

import (
	"context"
	"fmt"
	"strings"

	"github.com/pborman/uuid"
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

func (req *SignUpRequest) validate() error {
	if req.Username == "" {
		return fmt.Errorf("invalid: username is blank")
	}

	if req.Password == "" {
		return fmt.Errorf("invalid: password is blank")
	}

	return nil
}

type SignUpResponse struct {
	User *User
}

func (us *UserService) SignUp(ctx context.Context, req SignUpRequest) (*SignUpResponse, error) {
	req.normalize()

	if err := req.validate(); err != nil {
		return nil, fmt.Errorf("sign up failed: %w", err)
	}

	user := &User{
		ID:       uuid.New(),
		Username: req.Username,
	}

	resp := &SignUpResponse{User: user}
	return resp, nil
}

type User struct {
	ID       string
	Username string
}
