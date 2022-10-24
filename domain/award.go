package domain

import (
	"github.com/haleyrc/pkg/time"
)

type Award struct {
	CheevoID string

	UserID string

	Awarded time.Time
}

func (a *Award) Model() string { return "Award" }

func (a *Award) Validate() error {
	ve := NewValidationError(a)

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
