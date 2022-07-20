package cheevos

import (
	"context"
	"fmt"
	"strings"

	"github.com/pborman/uuid"
)

// UserService represents the main entrypoint for managing user.
type UserService struct {
	DB Database
}

// SignUpRequest represents the parameters for signing up a new user.
type SignUpRequest struct {
	// The username is required since it is used as the primary means of referring
	// to users within the app. This value is used for all display and search
	// purposes in lieu of an email address, which may be organization-specific
	// (or not, which may be worse) and is difficult to read.
	Username string

	// The password is required in order for the user to sign in. Passwords are
	// accepted as plaintext (and thus should only be transmitted over secure
	// channels for sign up), but are stored only in encrypted form.
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

// SignUpResponse is returned when a new user is successfully signed up.
type SignUpResponse struct {
	// The complete persisted user. The ID returned on the model is a unique
	// identifer for the user for use in further operations.
	User *User
}

// SignUp creates a new user and persists it to the database. It returns a
// response containing the new organization if successful.
func (us *UserService) SignUp(ctx context.Context, req SignUpRequest) (*SignUpResponse, error) {
	req.normalize()

	if err := req.validate(); err != nil {
		return nil, fmt.Errorf("sign up failed: %w", err)
	}

	user := &User{
		ID:       uuid.New(),
		Username: req.Username,
	}
	err := us.DB.Call(ctx, func(ctx context.Context, tx Transaction) error {
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("sign up failed: %w", err)
	}

	resp := &SignUpResponse{User: user}
	return resp, nil
}
