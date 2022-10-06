package service_test

import (
	"testing"

	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/internal/fake"
	"github.com/haleyrc/cheevos/internal/service"
	"github.com/haleyrc/cheevos/internal/testutil"
)

func TestMembershipValidationReturnsNilForAValidMembership(t *testing.T) {
	m := fake.Membership(uuid.New(), uuid.New())
	if err := m.Validate(); err != nil {
		t.Errorf("Expected validate to return nil, but got %v.", err)
	}
}

func TestMembershipValidationReturnsAnErrorForAnInvalidMembership(t *testing.T) {
	var m service.Membership
	testutil.RunValidationTests(t, &m, "validation failed: Membership is invalid", map[string]string{
		"OrganizationID": "Organization ID can't be blank.",
		"UserID":         "User ID can't be blank.",
		"Joined":         "Joined time can't be blank.",
	})
}
