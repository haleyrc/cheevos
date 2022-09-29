package roster_test

import (
	"testing"

	"github.com/haleyrc/cheevos/internal/fake"
	"github.com/haleyrc/cheevos/internal/testutil"
	"github.com/haleyrc/cheevos/roster"
	"github.com/pborman/uuid"
)

func TestMembershipValidationReturnsNilForAValidMembership(t *testing.T) {
	m := fake.Membership(uuid.New(), uuid.New())
	if err := m.Validate(); err != nil {
		t.Errorf("Expected validate to return nil, but got %v.", err)
	}
}

func TestMembershipValidationReturnsAnErrorForAnInvalidMembership(t *testing.T) {
	var m roster.Membership
	testutil.RunValidationTests(t, &m, "validation failed: Membership is invalid", map[string]string{
		"OrganizationID": "Organization ID can't be blank.",
		"UserID":         "User ID can't be blank.",
		"Joined":         "Joined time can't be blank.",
	})
}
