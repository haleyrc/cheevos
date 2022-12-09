package domain_test

import (
	"testing"

	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/domain"
	"github.com/haleyrc/cheevos/internal/assert"
	"github.com/haleyrc/cheevos/internal/fake"
	"github.com/haleyrc/cheevos/internal/testutil"
)

func TestAwardValidationReturnsNilForAValidAward(t *testing.T) {
	a := fake.Award(uuid.New(), uuid.New())
	assert.Error(t, a.Validate()).IsNil()
}

func TestAwardValidationReturnsAnErrorForAnInvalidAward(t *testing.T) {
	var a domain.Award
	testutil.RunValidationTests(t, &a, "validation failed: Award is invalid", map[string]string{
		"CheevoID": "Cheevo ID can't be blank.",
		"UserID":   "User ID can't be blank.",
		"Awarded":  "Awarded time can't be blank.",
	})
}
