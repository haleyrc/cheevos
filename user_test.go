package cheevos_test

import (
	"context"
	"testing"

	"github.com/haleyrc/cheevos"
	"github.com/haleyrc/cheevos/internal/mock"
)

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
