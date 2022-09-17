package mock

import (
	"context"
	"fmt"

	"github.com/haleyrc/cheevos/cheevo"
	"github.com/haleyrc/cheevos/lib/db"
)

type AwardCheevoToUserArgs struct {
	CheevoID    string
	RecipientID string
}

type CreateCheevoArgs struct {
	Cheevo *cheevo.Cheevo
}

type CheevoRepository struct {
	AwardCheevoToUserFn     func(ctx context.Context, tx db.Transaction, recipientID, cheevoID string) (*cheevo.Award, error)
	AwardCheevoToUserCalled struct {
		Count int
		With  AwardCheevoToUserArgs
	}

	CreateCheevoFn     func(ctx context.Context, tx db.Transaction, cheevo *cheevo.Cheevo) error
	CreateCheevoCalled struct {
		Count int
		With  CreateCheevoArgs
	}
}

func (cr *CheevoRepository) AwardCheevoToUser(ctx context.Context, tx db.Transaction, recipientID, cheevoID string) (*cheevo.Award, error) {
	if cr.AwardCheevoToUserFn == nil {
		return nil, mockMethodNotDefined("AwardCheevoToUser")
	}
	cr.AwardCheevoToUserCalled.Count++
	cr.AwardCheevoToUserCalled.With = AwardCheevoToUserArgs{CheevoID: cheevoID, RecipientID: recipientID}
	return cr.AwardCheevoToUserFn(ctx, tx, recipientID, cheevoID)
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
