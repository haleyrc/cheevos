package domain

import (
	"context"
)

type CheevosService interface {
	AwardCheevoToUser(ctx context.Context, recipientID, cheevoID string) error
	CreateCheevo(ctx context.Context, name, description, orgID string) (*Cheevo, error)
	GetCheevo(ctx context.Context, id string) (*Cheevo, error)
}
