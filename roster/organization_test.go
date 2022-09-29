package roster_test

import (
	"testing"

	"github.com/haleyrc/cheevos/internal/fake"
	"github.com/haleyrc/cheevos/internal/testutil"
	"github.com/haleyrc/cheevos/roster"
	"github.com/pborman/uuid"
)

func TestNormalizingAnOrganizationNormalizesName(t *testing.T) {
	subject := roster.Organization{Name: testutil.UnsafeString}
	subject.Normalize()
	if subject.Name != testutil.SafeString {
		t.Errorf("Expected organization name to be normalized, but it wasn't.")
	}
}

func TestOrganizationValidationReturnsNilForAValidOrganization(t *testing.T) {
	o := fake.Organization(uuid.New())
	if err := o.Validate(); err != nil {
		t.Errorf("Expected validate to return nil, but got %v.", err)
	}
}

func TestOrganizationValidationReturnsAnErrorForAnInvalidOrganization(t *testing.T) {
	var o roster.Organization
	testutil.RunValidationTests(t, &o, "validation failed: Organization is invalid", map[string]string{
		"ID":      "ID can't be blank.",
		"Name":    "Name can't be blank.",
		"OwnerID": "Owner ID can't be blank.",
	})
}
