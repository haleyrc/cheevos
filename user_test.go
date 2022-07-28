package cheevos_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/haleyrc/cheevos"
	"github.com/haleyrc/cheevos/internal/mock"
)

func TestUserLoggerLogsAnErrorFromSignUp(t *testing.T) {
	logger := NewTestLogger()

	cl := &cheevos.UserLogger{
		Svc: &mock.UserService{
			SignUpFn: func(_ context.Context, req cheevos.SignUpRequest) (*cheevos.SignUpResponse, error) {
				return nil, fmt.Errorf("oops")
			},
		},
		Logger: logger,
	}
	cl.SignUp(context.Background(), cheevos.SignUpRequest{
		Username: "Test",
		Password: "Testtest123",
	})

	logger.ShouldLog(t,
		`{"Fields":{"Username":"Test"},"Message":"signing up user"}`,
		`{"Fields":{"Error":"oops"},"Message":"sign up failed"}`,
	)
}

func TestUserLoggerLogsTheResponseFromSignUp(t *testing.T) {
	logger := NewTestLogger()

	cl := &cheevos.UserLogger{
		Svc: &mock.UserService{
			SignUpFn: func(_ context.Context, req cheevos.SignUpRequest) (*cheevos.SignUpResponse, error) {
				return &cheevos.SignUpResponse{
					User: &cheevos.User{
						ID:       "8059dcd7-bcc1-46fa-bfc0-3926c0b2c6ea",
						Username: "Test",
					},
				}, nil
			},
		},
		Logger: logger,
	}
	cl.SignUp(context.Background(), cheevos.SignUpRequest{
		Username: "Test",
		Password: "Testtest123",
	})

	logger.ShouldLog(t,
		`{"Fields":{"Username":"Test"},"Message":"signing up user"}`,
		`{"Fields":{"User":{"ID":"8059dcd7-bcc1-46fa-bfc0-3926c0b2c6ea","Username":"Test"}},"Message":"user signed up"}`,
	)
}

func TestCreatingAValidUserWithSucceeds(t *testing.T) {
	ctx := context.Background()
	db := mock.NewDatabase()
	svc := cheevos.UserService{DB: db}

	resp, err := svc.SignUp(ctx, cheevos.SignUpRequest{
		Username: "test",
		Password: "testtest",
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.User.ID == "" {
		t.Error("ID shouldn't be blank, but it was.")
	}
	if resp.User.Username != "test" {
		t.Errorf("Username should be \"test\", but got %q.", resp.User.Username)
	}
}

func TestCreatingAUserWithAnInvalidUsernameFails(t *testing.T) {
	ctx := context.Background()
	svc := cheevos.UserService{}

	testcases := map[string]string{
		"empty username": "",
		"blank username": " \t\n",
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			_, err := svc.SignUp(ctx, cheevos.SignUpRequest{
				Username: tc,
				Password: "testtest",
			})
			if err == nil {
				t.Error("expected an error, but got none")
			}
		})
	}
}

func TestCreatingAUserWithAnInvalidPasswordFails(t *testing.T) {
	ctx := context.Background()
	svc := cheevos.UserService{}

	testcases := map[string]string{
		"empty password": "",
		"blank password": " \t\n",
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			_, err := svc.SignUp(ctx, cheevos.SignUpRequest{
				Username: "test",
				Password: tc,
			})
			if err == nil {
				t.Error("expected an error, but got none")
			}
		})
	}
}
