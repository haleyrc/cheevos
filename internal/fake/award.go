package fake

import (
	"github.com/haleyrc/pkg/time"

	"github.com/haleyrc/cheevos/domain"
)

func Award(cheevoID, userID string) *domain.Award {
	return &domain.Award{
		CheevoID: cheevoID,
		UserID:   userID,
		Awarded:  time.Now(),
	}
}
