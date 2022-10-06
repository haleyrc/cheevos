package service_test

import (
	"testing"

	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/internal/fake"
	"github.com/haleyrc/cheevos/internal/service"
	"github.com/haleyrc/cheevos/internal/testutil"
)

func TestAwardValidationReturnsNilForAValidAward(t *testing.T) {
	a := fake.Award(uuid.New(), uuid.New())
	if err := a.Validate(); err != nil {
		t.Errorf("Expected validate to return nil, but got %v.", err)
	}
}

func TestAwardValidationReturnsAnErrorForAnInvalidAward(t *testing.T) {
	var a service.Award
	testutil.RunValidationTests(t, &a, "validation failed: Award is invalid", map[string]string{
		"CheevoID": "Cheevo ID can't be blank.",
		"UserID":   "User ID can't be blank.",
		"Awarded":  "Awarded time can't be blank.",
	})
}
