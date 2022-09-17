package mock

import (
	"context"

	"github.com/haleyrc/cheevos/award"
	"github.com/haleyrc/cheevos/lib/db"
)

type CreateAwardArgs struct {
	Award *award.Award
}

type AwardRepository struct {
	CreateAwardFn     func(ctx context.Context, tx db.Transaction, a *award.Award) error
	CreateAwardCalled struct {
		Count int
		With  CreateAwardArgs
	}
}

func (ar *AwardRepository) CreateAward(ctx context.Context, tx db.Transaction, a *award.Award) error {
	if ar.CreateAwardFn == nil {
		return mockMethodNotDefined("CreateAward")
	}
	ar.CreateAwardCalled.Count++
	ar.CreateAwardCalled.With = CreateAwardArgs{Award: a}
	return ar.CreateAwardFn(ctx, tx, a)
}
