package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/haleyrc/cheevos/domain"
	"github.com/haleyrc/cheevos/internal/mock"
	"github.com/haleyrc/cheevos/internal/password"
	"github.com/haleyrc/cheevos/internal/testutil"
)

func TestLoggerLogsAnErrorFromGetUser(t *testing.T) {
	logger := testutil.NewTestLogger()

	al := &authLogger{
		Service: &mock.AuthService{
			GetUserFn: func(_ context.Context, _ string) (*domain.User, error) {
				return nil, fmt.Errorf("oops")
			},
		},
		Logger: logger,
	}
	al.GetUser(context.Background(), "id")

	logger.ShouldLog(t,
		`{"Fields":{"ID":"id"},"Message":"getting user"}`,
		`{"Fields":{"Error":"oops"},"Message":"get user failed"}`,
	)
}

func TestLoggerLogsTheResponseFromGetUser(t *testing.T) {
	logger := testutil.NewTestLogger()

	al := &authLogger{
		Service: &mock.AuthService{
			GetUserFn: func(_ context.Context, _ string) (*domain.User, error) {
				return &domain.User{ID: "id", Username: "username"}, nil
			},
		},
		Logger: logger,
	}
	al.GetUser(context.Background(), "id")

	logger.ShouldLog(t,
		`{"Fields":{"ID":"id"},"Message":"getting user"}`,
		`{"Fields":{"User":{"ID":"id","Username":"username"}},"Message":"got user"}`,
	)

}

func TestLoggerLogsAnErrorFromSignUp(t *testing.T) {
	logger := testutil.NewTestLogger()

	al := &authLogger{
		Service: &mock.AuthService{
			SignUpFn: func(_ context.Context, _ string, _ password.Password) (*domain.User, error) {
				return nil, fmt.Errorf("oops")
			},
		},
		Logger: logger,
	}
	al.SignUp(context.Background(), "username", password.New("password"))

	logger.ShouldLog(t,
		`{"Fields":{"Username":"username"},"Message":"signing up user"}`,
		`{"Fields":{"Error":"oops"},"Message":"sign up failed"}`,
	)
}

func TestLoggerLogsTheResponseFromSignUp(t *testing.T) {
	logger := testutil.NewTestLogger()

	al := &authLogger{
		Service: &mock.AuthService{
			SignUpFn: func(_ context.Context, _ string, _ password.Password) (*domain.User, error) {
				return &domain.User{ID: "id", Username: "username"}, nil
			},
		},
		Logger: logger,
	}
	al.SignUp(context.Background(), "username", password.New("password"))

	logger.ShouldLog(t,
		`{"Fields":{"Username":"username"},"Message":"signing up user"}`,
		`{"Fields":{"User":{"ID":"id","Username":"username"}},"Message":"user signed up"}`,
	)
}
