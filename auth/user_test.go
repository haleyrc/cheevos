package auth_test

import (
	"testing"

	"github.com/haleyrc/cheevos/auth"
	"github.com/haleyrc/cheevos/internal/testutil"
)

func TestNormalizingAUserNormalizesUsername(t *testing.T) {
	subject := auth.User{Username: testutil.UnsafeString}
	subject.Normalize()
	if subject.Username != testutil.SafeString {
		t.Errorf("Expected user username to be normalized, but it wasn't.")
	}
}

func TestValidatingAUser(t *testing.T) {
	testcases := map[string]struct {
		input auth.User
		err   string
	}{
		"returns an error for a missing id": {
			input: auth.User{ID: "", Username: "username"},
			err:   "id is blank",
		},
		"returns an error for a missing username": {
			input: auth.User{ID: "id", Username: ""},
			err:   "username is blank",
		},
		"returns nil for a valid user": {
			input: auth.User{ID: "id", Username: "username"},
			err:   "",
		},
	}
	for name, tc := range testcases {
		testutil.RunValidationTests(t, name, &tc.input, tc.err)
	}
}
