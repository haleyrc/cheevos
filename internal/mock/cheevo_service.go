package mock

import (
	"context"

	"github.com/haleyrc/cheevos"
)

type CheevoService struct {
	AwardCheevoToUserFn func(context.Context, cheevos.AwardCheevoToUserRequest) (*cheevos.AwardCheevoToUserResponse, error)
	CreateCheevoFn      func(context.Context, cheevos.CreateCheevoRequest) (*cheevos.CreateCheevoResponse, error)
}

func (cs *CheevoService) AwardCheevoToUser(ctx context.Context, req cheevos.AwardCheevoToUserRequest) (*cheevos.AwardCheevoToUserResponse, error) {
	return cs.AwardCheevoToUserFn(ctx, req)
}

func (cs *CheevoService) CreateCheevo(ctx context.Context, req cheevos.CreateCheevoRequest) (*cheevos.CreateCheevoResponse, error) {
	return cs.CreateCheevoFn(ctx, req)
}
