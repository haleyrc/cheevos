package award

import (
	"context"
	"fmt"

	"github.com/haleyrc/cheevos/lib/db"
	"github.com/haleyrc/cheevos/lib/time"
)

type AwardRepository interface {
	CreateAward(ctx context.Context, tx db.Transaction, award *Award) error
}

type AwardService struct {
	DB   db.Database
	Repo AwardRepository
}

// AwardCheevoToUser awards a specific Cheevo to a User. Statistics for this
// event are bidirectional; a Cheevo "tracks" the number of Users that have
// received it and Users "track" how many Cheevos they have received.
func (as *AwardService) AwardCheevoToUser(ctx context.Context, recipientID, cheevoID string) error {
	err := as.DB.Call(ctx, func(ctx context.Context, tx db.Transaction) error {
		award := &Award{
			CheevoID: cheevoID,
			UserID:   recipientID,
			Awarded:  time.Now(),
		}
		if err := award.Validate(); err != nil {
			return err
		}
		return as.Repo.CreateAward(ctx, tx, award)
	})
	if err != nil {
		return fmt.Errorf("award cheevo to user failed: %w", err)
	}

	return nil
}
