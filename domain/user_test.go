package domain_test

import (
	"testing"

	"github.com/haleyrc/cheevos/domain"
	"github.com/haleyrc/cheevos/internal/assert"
	"github.com/haleyrc/cheevos/internal/fake"
	"github.com/haleyrc/cheevos/internal/testutil"
)

func TestNormalizingAUserNormalizesUsername(t *testing.T) {
	subject := domain.User{Username: testutil.UnsafeString}
	subject.Normalize()
	assert.String(t, "username", subject.Username).Equals(testutil.SafeString)
}

func TestUserValidationReturnsNilForAValidUser(t *testing.T) {
	u := fake.User()
	assert.Error(t, u.Validate()).IsNil()
}

func TestUserValidationReturnsAnErrorForAnInvalidUser(t *testing.T) {
	var u domain.User
	testutil.RunValidationTests(t, &u, "validation failed: User is invalid", map[string]string{
		"ID":       "ID can't be blank.",
		"Username": "Username can't be blank.",
	})
}
