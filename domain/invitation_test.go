package domain_test

import (
	"testing"

	"github.com/haleyrc/pkg/time"
	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/domain"
	"github.com/haleyrc/cheevos/internal/assert"
	"github.com/haleyrc/cheevos/internal/fake"
	"github.com/haleyrc/cheevos/internal/testutil"
)

func TestAnInvitationWithAnExpiresInThePastIsExpired(t *testing.T) {
	i := domain.Invitation{Expires: time.Now().Sub(time.Hour)}
	assert.Bool(t, "expired", i.Expired()).IsTrue()
}

func TestAnInvitationWithAnExpiresInTheFutureIsNotExpired(t *testing.T) {
	i := domain.Invitation{Expires: time.Now().Add(time.Hour)}
	assert.Bool(t, "expired", i.Expired()).IsFalse()
}

func TestNormalizingAnInvitationNormalizesEmail(t *testing.T) {
	subject := domain.Invitation{Email: testutil.UnsafeString}
	subject.Normalize()
	assert.String(t, "email", subject.Email).Equals(testutil.SafeString)
}

func TestInvitationValidationReturnsNilForAValidInvitation(t *testing.T) {
	i := fake.Invitation(uuid.New())
	assert.Error(t, i.Validate()).IsNil()
}

func TestInvitationValidationReturnsAnErrorForAnInvalidInvitation(t *testing.T) {
	var i domain.Invitation
	testutil.RunValidationTests(t, &i, "validation failed: Invitation is invalid", map[string]string{
		"ID":             "ID can't be blank.",
		"Email":          "Email can't be blank.",
		"OrganizationID": "Organization ID can't be blank.",
		"Expires":        "Expiration time can't be blank.",
	})
}
