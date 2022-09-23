package fake

import (
	"github.com/haleyrc/cheevos/cheevos"
	"github.com/haleyrc/cheevos/lib/time"
)

func Award(cheevoID, userID string) *cheevos.Award {
	return &cheevos.Award{
		CheevoID: cheevoID,
		UserID:   userID,
		Awarded:  time.Now(),
	}
}
