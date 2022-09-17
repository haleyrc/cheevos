package user_test

import (
	"strings"
	"testing"

	"github.com/haleyrc/cheevos/internal/testutil"
	"github.com/haleyrc/cheevos/user"
)

func TestNormalizingAUserNormalizesUsername(t *testing.T) {
	subject := user.User{Username: testutil.UnsafeString}
	subject.Normalize()
	if subject.Username != testutil.SafeString {
		t.Errorf("Expected user username to be normalized, but it wasn't.")
	}
}

func TestValidatingAUser(t *testing.T) {
	testcases := map[string]struct {
		input user.User
		err   string
	}{
		"returns an error for a missing id": {
			input: user.User{ID: "", Username: "username"},
			err:   "id is blank",
		},
		"returns an error for a missing username": {
			input: user.User{ID: "id", Username: ""},
			err:   "username is blank",
		},
		"returns nil for a valid user": {
			input: user.User{ID: "id", Username: "username"},
			err:   "",
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			err := tc.input.Validate()
			if tc.err == "" {
				if err != nil {
					t.Error("unexpected error:", err)
				}
			} else {
				if err == nil {
					t.Errorf("expected error, but got nil")
					return
				}
				compareError(t, tc.err, err)
			}
		})
	}
}

func compareError(t *testing.T, want string, got error) {
	t.Helper()
	g := got.Error()
	if !strings.Contains(g, want) {
		t.Errorf("error %q does not include %q", g, want)
	}
}
