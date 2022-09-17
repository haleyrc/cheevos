package mock

import (
	"context"

	"github.com/haleyrc/cheevos/cheevo"
)

type CheevoService struct {
	CreateCheevoFn func(ctx context.Context, name, description, orgID string) (*cheevo.Cheevo, error)
}

func (cs *CheevoService) CreateCheevo(ctx context.Context, name, description, orgID string) (*cheevo.Cheevo, error) {
	return cs.CreateCheevoFn(ctx, name, description, orgID)
}
