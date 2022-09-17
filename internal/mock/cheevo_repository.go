package mock

import (
	"context"
	"fmt"

	"github.com/haleyrc/cheevos/cheevo"
	"github.com/haleyrc/cheevos/lib/db"
)

type CreateCheevoArgs struct {
	Cheevo *cheevo.Cheevo
}

type CheevoRepository struct {
	CreateCheevoFn     func(ctx context.Context, tx db.Transaction, cheevo *cheevo.Cheevo) error
	CreateCheevoCalled struct {
		Count int
		With  CreateCheevoArgs
	}
}

func (cr *CheevoRepository) CreateCheevo(ctx context.Context, tx db.Transaction, cheevo *cheevo.Cheevo) error {
	if cr.CreateCheevoFn == nil {
		return mockMethodNotDefined("CreateCheevo")
	}
	cr.CreateCheevoCalled.Count++
	cr.CreateCheevoCalled.With = CreateCheevoArgs{Cheevo: cheevo}
	return cr.CreateCheevoFn(ctx, tx, cheevo)
}

func mockMethodNotDefined(funcName string) error {
	return fmt.Errorf("mock method %s is not defined", funcName)
}
