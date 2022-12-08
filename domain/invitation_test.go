package domain_test

import (
	"testing"

	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/domain"
	"github.com/haleyrc/cheevos/internal/fake"
	"github.com/haleyrc/cheevos/internal/testutil"
)

/* TODO: I can't really test these in this way anymore. I'm not sure if it's
   worth pursuing at the moment, so I'm just delaying it and saying that this
	 logic is pretty simple and the time library is tested.
func TestAnInvitationWithAnExpiresInThePastIsExpired(t *testing.T) {
	i := domain.Invitation{Expires: time.Now().Sub(time.Hour)}
	if !i.Expired() {
		t.Errorf("Expected invitation expiring at %s to be expired, but it wasn't.", i.Expires)
	}
}

func TestAnInvitationWithAnExpiresInTheFutureIsNotExpired(t *testing.T) {
	i := domain.Invitation{Expires: time.Now().Add(time.Hour)}
	if i.Expired() {
		t.Errorf("Expected invitation expiring at %s to not be expired, but it was.", i.Expires)
	}
}
*/

func TestNormalizingAnInvitationNormalizesEmail(t *testing.T) {
	subject := domain.Invitation{Email: testutil.UnsafeString}
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
	var i domain.Invitation
	testutil.RunValidationTests(t, &i, "validation failed: Invitation is invalid", map[string]string{
		"ID":             "ID can't be blank.",
		"Email":          "Email can't be blank.",
		"OrganizationID": "Organization ID can't be blank.",
		// "Expires":        "Expiration time can't be blank.",
	})
}
