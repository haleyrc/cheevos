package cheevos_test

import (
	"context"
	"testing"

	"github.com/haleyrc/cheevos"
)

func TestCreatingAValidUserWithSucceeds(t *testing.T) {
	ctx := context.Background()
	svc := cheevos.UserService{}
	_, err := svc.SignUp(ctx, cheevos.SignUpRequest{
		Username: "test",
		Password: "testtest",
	})
	if err != nil {
		t.Fatal(err)
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
