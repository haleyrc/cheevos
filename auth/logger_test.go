package auth_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/haleyrc/cheevos/auth"
	"github.com/haleyrc/cheevos/internal/mock"
	"github.com/haleyrc/cheevos/internal/testutil"
)

func TestLoggerLogsAnErrorFromSignUp(t *testing.T) {
	logger := testutil.NewTestLogger()

	cl := &auth.Logger{
		Service: &mock.AuthService{
			SignUpFn: func(_ context.Context, _, _ string) (*auth.User, error) {
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

	cl := &auth.Logger{
		Service: &mock.AuthService{
			SignUpFn: func(_ context.Context, _, _ string) (*auth.User, error) {
				return &auth.User{ID: "id", Username: "username"}, nil
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
