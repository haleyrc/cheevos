package mock

import (
	"context"

	"github.com/haleyrc/cheevos/domain"
)

var _ domain.CheevosService = &CheevosService{}

type CheevosService struct {
	AwardCheevoToUserFn func(ctx context.Context, recipientID, cheevoID string) error
	CreateCheevoFn      func(ctx context.Context, name, description, orgID string) (*domain.Cheevo, error)
	GetCheevoFn         func(ctx context.Context, id string) (*domain.Cheevo, error)
}

func (cs *CheevosService) AwardCheevoToUser(ctx context.Context, recipientID, cheevoID string) error {
	return cs.AwardCheevoToUserFn(ctx, recipientID, cheevoID)
}

func (cs *CheevosService) CreateCheevo(ctx context.Context, name, description, orgID string) (*domain.Cheevo, error) {
	return cs.CreateCheevoFn(ctx, name, description, orgID)
}

func (cs *CheevosService) GetCheevo(ctx context.Context, id string) (*domain.Cheevo, error) {
	return cs.GetCheevoFn(ctx, id)
}
