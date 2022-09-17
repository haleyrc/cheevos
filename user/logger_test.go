package user_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/haleyrc/cheevos/internal/mock"
	"github.com/haleyrc/cheevos/internal/testutil"
	"github.com/haleyrc/cheevos/user"
)

func TestUserLoggerLogsAnErrorFromSignUp(t *testing.T) {
	logger := testutil.NewTestLogger()

	cl := &user.UserLogger{
		Svc: &mock.UserService{
			SignUpFn: func(_ context.Context, _, _ string) (*user.User, error) {
				return nil, fmt.Errorf("oops")
			},
		},
		Logger: logger,
	}
	cl.SignUp(context.Background(), "Test", "Testtest123")

	logger.ShouldLog(t,
		`{"Fields":{"Username":"Test"},"Message":"signing up user"}`,
		`{"Fields":{"Error":"oops"},"Message":"sign up failed"}`,
	)
}

func TestUserLoggerLogsTheResponseFromSignUp(t *testing.T) {
	logger := testutil.NewTestLogger()

	cl := &user.UserLogger{
		Svc: &mock.UserService{
			SignUpFn: func(_ context.Context, _, _ string) (*user.User, error) {
				return &user.User{
					ID:       "8059dcd7-bcc1-46fa-bfc0-3926c0b2c6ea",
					Username: "Test",
				}, nil
			},
		},
		Logger: logger,
	}
	cl.SignUp(context.Background(), "Test", "Testtest123")

	logger.ShouldLog(t,
		`{"Fields":{"Username":"Test"},"Message":"signing up user"}`,
		`{"Fields":{"User":{"ID":"8059dcd7-bcc1-46fa-bfc0-3926c0b2c6ea","Username":"Test"}},"Message":"user signed up"}`,
	)
}
