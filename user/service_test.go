package user_test

import (
	"context"
	"testing"

	"github.com/haleyrc/cheevos/internal/mock"
	"github.com/haleyrc/cheevos/lib/db"
	"github.com/haleyrc/cheevos/user"
)

func TestSigningUpSucceeds(t *testing.T) {
	ctx := context.Background()
	repo := mock.Repository{
		CreateUserFn: func(_ context.Context, _ db.Transaction, _ *user.User, _ string) error { return nil },
	}
	svc := user.UserService{
		DB:   &mock.Database{},
		Repo: &repo,
	}

	user, err := svc.SignUp(ctx, "test", "testtest")
	if err != nil {
		t.Fatal(err)
	}

	if repo.CreateUserCalled.Count != 1 {
		t.Errorf("Expected repository to receive CreateUser, but it didn't.")
	}
	if repo.CreateUserCalled.With.User.ID != user.ID {
		t.Errorf(
			"Expected repository.CreateUser to receive id %q, but got %q.",
			user.ID, repo.CreateUserCalled.With.User.ID,
		)
	}
	if repo.CreateUserCalled.With.User.Username != "test" {
		t.Errorf(
			"Expected repository.CreateUser to receive username %q, but got %q.",
			"test", repo.CreateUserCalled.With.User.Username,
		)
	}
	if repo.CreateUserCalled.With.PasswordHash == "testtest" {
		t.Errorf("Expected repository.CreateUser not to receive a plaintext password, but it did.")
	}

	if user.ID == "" {
		t.Error("ID shouldn't be blank, but it was.")
	}
	if user.Username != "test" {
		t.Errorf("Username should be \"test\", but got %q.", user.Username)
	}
}
