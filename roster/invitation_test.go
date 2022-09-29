package roster_test

import (
	"testing"

	"github.com/haleyrc/cheevos/internal/fake"
	"github.com/haleyrc/cheevos/internal/testutil"
	"github.com/haleyrc/cheevos/roster"
	"github.com/pborman/uuid"
)

func TestNormalizingAnInvitationNormalizesEmail(t *testing.T) {
	subject := roster.Invitation{Email: testutil.UnsafeString}
	subject.Normalize()
	if subject.Email != testutil.SafeString {
		t.Errorf("Expected invitation email to be normalized, but it wasn't.")
	}
}

func TestInvitationValidationReturnsNilForAValidInvitation(t *testing.T) {
	i := fake.Invitation(uuid.New())
	if err := i.Validate(); err != nil {
		t.Errorf("Expected validate to return nil, but got %v.", err)
	}
}

func TestInvitationValidationReturnsAnErrorForAnInvalidInvitation(t *testing.T) {
	var i roster.Invitation
	testutil.RunValidationTests(t, &i, "validation failed: Invitation is invalid", map[string]string{
		"ID":             "ID can't be blank.",
		"Email":          "Email can't be blank.",
		"OrganizationID": "Organization ID can't be blank.",
		"Expires":        "Expiration time can't be blank.",
	})
}
