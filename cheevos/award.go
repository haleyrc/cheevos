package cheevos

import (
	"github.com/haleyrc/cheevos/core"
	"github.com/haleyrc/cheevos/lib/time"
)

type Award struct {
	CheevoID string

	UserID string

	Awarded time.Time
}

func (a *Award) Model() string { return "Award" }

func (a *Award) Validate() error {
	ve := core.NewValidationError(a)

	if a.CheevoID == "" {
		ve.Add("CheevoID", "Cheevo ID can't be blank.")
	}

	if a.UserID == "" {
		ve.Add("UserID", "User ID can't be blank.")
	}

	if a.Awarded.IsZero() {
		ve.Add("Awarded", "Awarded time can't be blank.")
	}

	return ve.Error()
}
