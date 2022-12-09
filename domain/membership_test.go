package domain_test

import (
	"testing"

	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/domain"
	"github.com/haleyrc/cheevos/internal/assert"
	"github.com/haleyrc/cheevos/internal/fake"
	"github.com/haleyrc/cheevos/internal/testutil"
)

func TestMembershipValidationReturnsNilForAValidMembership(t *testing.T) {
	m := fake.Membership(uuid.New(), uuid.New())
	assert.Error(t, m.Validate()).IsNil()
}

func TestMembershipValidationReturnsAnErrorForAnInvalidMembership(t *testing.T) {
	var m domain.Membership
	testutil.RunValidationTests(t, &m, "validation failed: Membership is invalid", map[string]string{
		"OrganizationID": "Organization ID can't be blank.",
		"UserID":         "User ID can't be blank.",
		"Joined":         "Joined time can't be blank.",
	})
}
