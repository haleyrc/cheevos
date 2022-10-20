package fake

import (
	"github.com/haleyrc/pkg/time"

	"github.com/haleyrc/cheevos"
)

func Award(cheevoID, userID string) *cheevos.Award {
	return &cheevos.Award{
		CheevoID: cheevoID,
		UserID:   userID,
		Awarded:  time.Now(),
	}
}
