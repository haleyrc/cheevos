package mock

import (
	"context"

	"github.com/haleyrc/cheevos/cheevo"
)

type CheevoService struct {
	AwardCheevoToUserFn func(ctx context.Context, userID, cheevoID string) error
	CreateCheevoFn      func(ctx context.Context, name, description, orgID string) (*cheevo.Cheevo, error)
}

func (cs *CheevoService) AwardCheevoToUser(ctx context.Context, userID, cheevoID string) error {
	return cs.AwardCheevoToUserFn(ctx, userID, cheevoID)
}

func (cs *CheevoService) CreateCheevo(ctx context.Context, name, description, orgID string) (*cheevo.Cheevo, error) {
	return cs.CreateCheevoFn(ctx, name, description, orgID)
}
