package service

import (
	"context"
	"testing"

	"github.com/haleyrc/pkg/pg"
	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/domain"
	"github.com/haleyrc/cheevos/internal/assert"
	"github.com/haleyrc/cheevos/internal/mock"
	"github.com/haleyrc/cheevos/internal/password"
)

func TestGettingAUserSucceeds(t *testing.T) {
	var (
		assert = assert.New(t)
		ctx    = context.Background()
		mockDB = &mock.Database{}
		repo   = &mock.Repository{
			GetUserFn: func(_ context.Context, _ pg.Tx, _ *domain.User, _ string) error { return nil },
		}
		svc = &authService{DB: mockDB, Repo: repo}

		id = uuid.New()
	)

	_, err := svc.GetUser(ctx, id)
	assert.Error(err).IsUnexpected()
	assert.Int("calls to GetUser", repo.GetUserCalled.Count).Equals(1)
	assert.String("id", repo.GetUserCalled.With.ID).Equals(id)
}

func TestSigningUpSucceeds(t *testing.T) {
	var (
		assert = assert.New(t)
		ctx    = context.Background()
		mockDB = &mock.Database{}
		repo   = &mock.Repository{
			InsertUserFn: func(_ context.Context, _ pg.Tx, _ *domain.User, _ string) error { return nil },
		}
		svc = &authService{
			DB:                mockDB,
			Repo:              repo,
			PasswordValidator: password.NewValidator(),
		}

		username = "username"
		pw       = "password"
	)

	user, err := svc.SignUp(ctx, username, password.New(pw))

	assert.Error(err).IsUnexpected()
	assert.Int("calls to InsertUser", repo.InsertUserCalled.Count).Equals(1)
	assert.String("id", repo.InsertUserCalled.With.User.ID).Equals(user.ID)
	assert.String("username", repo.InsertUserCalled.With.User.Username).Equals(user.Username)
	assert.String("password hash", repo.InsertUserCalled.With.PasswordHash).NotEquals(pw)
	assert.String("id", user.ID).NotBlank()
	assert.String("username", user.Username).Equals(username)
}
