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
	var fes []FieldError

	if a.CheevoID == "" {
		fes = append(fes, FieldError{
			Field: "CheevoID", Msg: "Cheevo ID can't be blank.",
		})
	}

	if a.UserID == "" {
		fes = append(fes, FieldError{
			Field: "UserID", Msg: "User ID can't be blank.",
		})
	}

	if a.Awarded.IsZero() {
		fes = append(fes, FieldError{
			Field: "Awarded", Msg: "Awarded time can't be blank.",
		})
	}

	if len(fes) > 0 {
		return NewValidationError("Award", fes)
	}

	return nil
}
