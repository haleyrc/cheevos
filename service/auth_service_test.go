package service

import (
	"context"
	"testing"

	"github.com/haleyrc/pkg/pg"
	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/domain"
	"github.com/haleyrc/cheevos/internal/mock"
)

func TestGettingAUserSucceeds(t *testing.T) {
	var (
		ctx = context.Background()

		id = uuid.New()

		mockDB = &mock.Database{}

		repo = &mock.Repository{
			GetUserFn: func(_ context.Context, _ pg.Tx, _ *domain.User, _ string) error { return nil },
		}

		svc = &authService{
			DB:   mockDB,
			Repo: repo,
		}
	)

	_, err := svc.GetUser(ctx, id)
	if err != nil {
		t.Fatal(err)
	}

	if repo.GetUserCalled.Count != 1 {
		t.Errorf("Expected repository to receive GetUser, but it didn't.")
	}
	if repo.GetUserCalled.With.ID != id {
		t.Errorf(
			"Expected repository.GetUser to receive id %q, but got %q.",
			id, repo.GetUserCalled.With.ID,
		)
	}
}

func TestSigningUpSucceeds(t *testing.T) {
	var (
		ctx = context.Background()

		username = "username"
		password = "password"

		mockDB = &mock.Database{}

		repo = &mock.Repository{
			InsertUserFn: func(_ context.Context, _ pg.Tx, _ *domain.User, _ string) error { return nil },
		}

		svc = &authService{
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
