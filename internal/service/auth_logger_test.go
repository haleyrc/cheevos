package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/haleyrc/cheevos"
	"github.com/haleyrc/cheevos/internal/mock"
	"github.com/haleyrc/cheevos/internal/testutil"
)

func TestLoggerLogsAnErrorFromSignUp(t *testing.T) {
	logger := testutil.NewTestLogger()

	cl := &authLogger{
		Service: &mock.AuthService{
			SignUpFn: func(_ context.Context, _, _ string) (*cheevos.User, error) {
				return nil, fmt.Errorf("oops")
			},
		},
		Logger: logger,
	}
	cl.SignUp(context.Background(), "username", "password")

	logger.ShouldLog(t,
		`{"Fields":{"Username":"username"},"Message":"signing up user"}`,
		`{"Fields":{"Error":"oops"},"Message":"sign up failed"}`,
	)
}

func TestLoggerLogsTheResponseFromSignUp(t *testing.T) {
	logger := testutil.NewTestLogger()

	cl := &authLogger{
		Service: &mock.AuthService{
			SignUpFn: func(_ context.Context, _, _ string) (*cheevos.User, error) {
				return &cheevos.User{ID: "id", Username: "username"}, nil
			},
		},
		Logger: logger,
	}
	cl.SignUp(context.Background(), "username", "password")

	logger.ShouldLog(t,
		`{"Fields":{"Username":"username"},"Message":"signing up user"}`,
		`{"Fields":{"User":{"ID":"id","Username":"username"}},"Message":"user signed up"}`,
	)
}
