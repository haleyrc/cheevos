package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/haleyrc/cheevos/internal/mock"
	"github.com/haleyrc/cheevos/internal/service"
	"github.com/haleyrc/cheevos/internal/testutil"
)

func TestLoggerLogsAnErrorFromSignUp(t *testing.T) {
	logger := testutil.NewTestLogger()

	cl := &service.AuthLogger{
		Service: &mock.AuthService{
			SignUpFn: func(_ context.Context, _, _ string) (*service.User, error) {
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

	cl := &service.AuthLogger{
		Service: &mock.AuthService{
			SignUpFn: func(_ context.Context, _, _ string) (*service.User, error) {
				return &service.User{ID: "id", Username: "username"}, nil
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
