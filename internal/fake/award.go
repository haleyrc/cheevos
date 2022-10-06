package fake

import (
	"github.com/haleyrc/cheevos/internal/lib/time"
	"github.com/haleyrc/cheevos/internal/service"
)

func Award(cheevoID, userID string) *service.Award {
	return &service.Award{
		CheevoID: cheevoID,
		UserID:   userID,
		Awarded:  time.Now(),
	}
}
