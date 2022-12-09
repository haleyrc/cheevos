package domain_test

import (
	"testing"

	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/domain"
	"github.com/haleyrc/cheevos/internal/assert"
	"github.com/haleyrc/cheevos/internal/fake"
	"github.com/haleyrc/cheevos/internal/testutil"
)

func TestNormalizingAnOrganizationNormalizesName(t *testing.T) {
	subject := domain.Organization{Name: testutil.UnsafeString}
	subject.Normalize()
	assert.String(t, "organization name", subject.Name).Equals(testutil.SafeString)
}

func TestOrganizationValidationReturnsNilForAValidOrganization(t *testing.T) {
	o := fake.Organization(uuid.New())
	assert.Error(t, o.Validate()).IsNil()
}

func TestOrganizationValidationReturnsAnErrorForAnInvalidOrganization(t *testing.T) {
	var o domain.Organization
	testutil.RunValidationTests(t, &o, "validation failed: Organization is invalid", map[string]string{
		"ID":      "ID can't be blank.",
		"Name":    "Name can't be blank.",
		"OwnerID": "Owner ID can't be blank.",
	})
}
