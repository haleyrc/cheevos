package mock

import (
	"context"

	"github.com/haleyrc/cheevos/internal/server"
	"github.com/haleyrc/cheevos/internal/service"
)

var _ server.CheevosService = &CheevosService{}

type CheevosService struct {
	AwardCheevoToUserFn func(ctx context.Context, recipientID, cheevoID string) error
	CreateCheevoFn      func(ctx context.Context, name, description, orgID string) (*service.Cheevo, error)
	GetCheevoFn         func(ctx context.Context, id string) (*service.Cheevo, error)
}

func (cs *CheevosService) AwardCheevoToUser(ctx context.Context, recipientID, cheevoID string) error {
	return cs.AwardCheevoToUserFn(ctx, recipientID, cheevoID)
}

func (cs *CheevosService) CreateCheevo(ctx context.Context, name, description, orgID string) (*service.Cheevo, error) {
	return cs.CreateCheevoFn(ctx, name, description, orgID)
}

func (cs *CheevosService) GetCheevo(ctx context.Context, id string) (*service.Cheevo, error) {
	return cs.GetCheevoFn(ctx, id)
}
