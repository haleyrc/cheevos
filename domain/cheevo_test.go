package domain_test

import (
	"testing"

	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/domain"
	"github.com/haleyrc/cheevos/internal/assert"
	"github.com/haleyrc/cheevos/internal/fake"
	"github.com/haleyrc/cheevos/internal/testutil"
)

func TestNormalizingACheevoNormalizesName(t *testing.T) {
	subject := domain.Cheevo{Name: testutil.UnsafeString}
	subject.Normalize()
	assert.String(t, "cheevo name", subject.Name).Equals(testutil.SafeString)
}

func TestNormalizingACheevoNormalizesDescription(t *testing.T) {
	subject := domain.Cheevo{Description: testutil.UnsafeString}
	subject.Normalize()
	assert.String(t, "cheevo description", subject.Description).Equals(testutil.SafeString)
}

func TestCheevoValidationReturnsNilForAValidCheevo(t *testing.T) {
	c := fake.Cheevo(uuid.New())
	assert.Error(t, c.Validate()).IsNil()
}

func TestCheevoValidationReturnsAnErrorForAnInvalidCheevo(t *testing.T) {
	var c domain.Cheevo
	testutil.RunValidationTests(t, &c, "validation failed: Cheevo is invalid", map[string]string{
		"ID":             "ID can't be blank.",
		"Name":           "Name can't be blank.",
		"Description":    "Description can't be blank.",
		"OrganizationID": "Organization ID can't be blank.",
	})
}
