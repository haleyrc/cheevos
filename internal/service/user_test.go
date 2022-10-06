package service_test

import (
	"testing"

	"github.com/haleyrc/cheevos/auth"
	"github.com/haleyrc/cheevos/internal/fake"
	"github.com/haleyrc/cheevos/internal/testutil"
)

func TestNormalizingAUserNormalizesUsername(t *testing.T) {
	subject := auth.User{Username: testutil.UnsafeString}
	subject.Normalize()
	if subject.Username != testutil.SafeString {
		t.Errorf("Expected user username to be normalized, but it wasn't.")
	}
}

func TestUserValidationReturnsNilForAValidUser(t *testing.T) {
	u := fake.User()
	if err := u.Validate(); err != nil {
		t.Errorf("Expected validate to return nil, but got %v.", err)
	}
}

func TestUserValidationReturnsAnErrorForAnInvalidUser(t *testing.T) {
	var u auth.User
	testutil.RunValidationTests(t, &u, "validation failed: User is invalid", map[string]string{
		"ID":       "ID can't be blank.",
		"Username": "Username can't be blank.",
	})
}
