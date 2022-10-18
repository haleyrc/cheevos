package fake

import (
	"github.com/haleyrc/cheevos"
	"github.com/haleyrc/cheevos/internal/lib/time"
)

func Award(cheevoID, userID string) *cheevos.Award {
	return &cheevos.Award{
		CheevoID: cheevoID,
		UserID:   userID,
		Awarded:  time.Now(),
	}
}
