package service

import (
	"context"
	"testing"

	"github.com/haleyrc/pkg/pg"
	"github.com/haleyrc/pkg/time"
	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/domain"
	"github.com/haleyrc/cheevos/internal/assert"
	"github.com/haleyrc/cheevos/internal/mock"
)

func TestAwardingACheevoSucceeds(t *testing.T) {
	var (
		assert = assert.New(t)
		ctx    = context.Background()
		mockDB = &mock.Database{}
		repo   = &mock.Repository{
			InsertAwardFn: func(_ context.Context, _ pg.Tx, _ *domain.Award) error { return nil },
		}
		svc = &cheevosService{DB: mockDB, Repo: repo}

		cheevoID = uuid.New()
		userID   = uuid.New()
		now      = time.Now()
	)

	err := svc.AwardCheevoToUser(ctx, userID, cheevoID)

	assert.Error(err).IsUnexpected()
	assert.Int("calls to InsertAward", repo.InsertAwardCalled.Count).Equals(1)
	assert.String("cheevo id", repo.InsertAwardCalled.With.Award.CheevoID).Equals(cheevoID)
	assert.String("user id", repo.InsertAwardCalled.With.Award.UserID).Equals(userID)
	assert.
		String("awarded", repo.InsertAwardCalled.With.Award.Awarded.String()).
		Equals(now.UTC().String())
}

func TestCreatingAValidCheevoSucceeds(t *testing.T) {
	var (
		assert = assert.New(t)
		ctx    = context.Background()
		mockDB = &mock.Database{}
		repo   = &mock.Repository{
			InsertCheevoFn: func(_ context.Context, _ pg.Tx, _ *domain.Cheevo) error { return nil },
		}
		svc = &cheevosService{DB: mockDB, Repo: repo}

		name        = "name"
		description = "description"
		orgID       = uuid.New()
	)

	cheevo, err := svc.CreateCheevo(ctx, name, description, orgID)
	assert.Error(err).IsUnexpected()

	assert.Int("calls to InsertCheevo", repo.InsertCheevoCalled.Count).Equals(1)
	assert.String("id", repo.InsertCheevoCalled.With.Cheevo.ID).Equals(cheevo.ID)
	assert.String("name", repo.InsertCheevoCalled.With.Cheevo.Name).Equals(name)
	assert.String("description", repo.InsertCheevoCalled.With.Cheevo.Description).Equals(description)

	assert.String("id", cheevo.ID).NotBlank()
	assert.String("name", cheevo.Name).Equals(name)
	assert.String("description", cheevo.Description).Equals(description)
}

func TestGettingACheevoSucceeds(t *testing.T) {
	var (
		assert = assert.New(t)
		ctx    = context.Background()
		mockDB = &mock.Database{}
		repo   = &mock.Repository{
			GetCheevoFn: func(_ context.Context, _ pg.Tx, _ *domain.Cheevo, _ string) error { return nil },
		}
		svc = &cheevosService{DB: mockDB, Repo: repo}

		id = uuid.New()
	)

	_, err := svc.GetCheevo(ctx, id)

	assert.Error(err).IsUnexpected()
	assert.Int("calls to GetCheevo", repo.GetCheevoCalled.Count).Equals(1)
	assert.String("id", repo.GetCheevoCalled.With.ID).Equals(id)
}
