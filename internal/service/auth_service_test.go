package service_test

import (
	"context"
	"testing"

	"github.com/haleyrc/cheevos/internal/lib/db"
	"github.com/haleyrc/cheevos/internal/mock"
	"github.com/haleyrc/cheevos/internal/service"
)

func TestSigningUpSucceeds(t *testing.T) {
	var (
		ctx = context.Background()

		username = "username"
		password = "password"

		mockDB = &mock.Database{}

		repo = &mock.Repository{
			InsertUserFn: func(_ context.Context, _ db.Tx, _ *service.User, _ string) error { return nil },
		}

		svc = service.AuthService{
			DB:   mockDB,
			Repo: repo,
		}
	)

	user, err := svc.SignUp(ctx, username, password)
	if err != nil {
		t.Fatal(err)
	}

	if repo.InsertUserCalled.Count != 1 {
		t.Errorf("Expected repository to receive InsertUser, but it didn't.")
	}
	if repo.InsertUserCalled.With.User.ID != user.ID {
		t.Errorf(
			"Expected repository.InsertUser to receive id %q, but got %q.",
			user.ID, repo.InsertUserCalled.With.User.ID,
		)
	}
	if repo.InsertUserCalled.With.User.Username != username {
		t.Errorf(
			"Expected repository.InsertUser to receive username %q, but got %q.",
			username, repo.InsertUserCalled.With.User.Username,
		)
	}
	if repo.InsertUserCalled.With.PasswordHash == password {
		t.Errorf("Expected repository.InsertUser not to receive a plaintext password, but it did.")
	}

	if user.ID == "" {
		t.Error("ID shouldn't be blank, but it was.")
	}
	if user.Username != username {
		t.Errorf("Username should be %q, but got %q.", username, user.Username)
	}
}
