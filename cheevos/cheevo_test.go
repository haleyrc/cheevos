package cheevos_test

import (
	"testing"

	"github.com/haleyrc/cheevos/cheevos"
	"github.com/haleyrc/cheevos/internal/fake"
	"github.com/haleyrc/cheevos/internal/testutil"
	"github.com/pborman/uuid"
)

func TestNormalizingACheevoNormalizesName(t *testing.T) {
	subject := cheevos.Cheevo{Name: testutil.UnsafeString}
	subject.Normalize()
	if subject.Name != testutil.SafeString {
		t.Errorf("Expected cheevo name to be normalized, but it wasn't.")
	}
}

func TestNormalizingACheevoNormalizesDescription(t *testing.T) {
	subject := cheevos.Cheevo{Description: testutil.UnsafeString}
	subject.Normalize()
	if subject.Description != testutil.SafeString {
		t.Errorf("Expected cheevo description to be normalized, but it wasn't.")
	}
}

func TestCheevoValidationReturnsNilForAValidCheevo(t *testing.T) {
	c := fake.Cheevo(uuid.New())
	if err := c.Validate(); err != nil {
		t.Errorf("Expected validate to return nil, but got %v.", err)
	}
}

func TestCheevoValidationReturnsAnErrorForAnInvalidCheevo(t *testing.T) {
	var c cheevos.Cheevo
	testutil.RunValidationTests(t, &c, "validation failed: Cheevo is invalid", map[string]string{
		"ID":             "ID can't be blank.",
		"Name":           "Name can't be blank.",
		"Description":    "Description can't be blank.",
		"OrganizationID": "Organization ID can't be blank.",
	})
}
