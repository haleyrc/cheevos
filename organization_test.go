package cheevos_test

import (
	"testing"

	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos"
	"github.com/haleyrc/cheevos/internal/fake"
	"github.com/haleyrc/cheevos/internal/testutil"
)

func TestNormalizingAnOrganizationNormalizesName(t *testing.T) {
	subject := cheevos.Organization{Name: testutil.UnsafeString}
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
	var o cheevos.Organization
	testutil.RunValidationTests(t, &o, "validation failed: Organization is invalid", map[string]string{
		"ID":      "ID can't be blank.",
		"Name":    "Name can't be blank.",
		"OwnerID": "Owner ID can't be blank.",
	})
}