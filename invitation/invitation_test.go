package invitation_test

import (
	"strings"
	"testing"
	"time"

	"github.com/haleyrc/cheevos/internal/testutil"
	"github.com/haleyrc/cheevos/invitation"
)

func TestNormalizingAnInvitationNormalizesEmail(t *testing.T) {
	subject := invitation.Invitation{Email: testutil.UnsafeString}
	subject.Normalize()
	if subject.Email != testutil.SafeString {
		t.Errorf("Expected invitation email to be normalized, but it wasn't.")
	}
}

func TestValidatingAnInvitation(t *testing.T) {
	testcases := map[string]struct {
		input invitation.Invitation
		err   string
	}{
		"returns an error for a missing email": {
			input: invitation.Invitation{Email: "", CodeHash: "codehash", Expires: time.Now()},
			err:   "email is blank",
		},
		"returns an error for a missing code hash": {
			input: invitation.Invitation{Email: "email", CodeHash: "", Expires: time.Now()},
			err:   "code hash is blank",
		},
		"returns an error for a missing expiration": {
			input: invitation.Invitation{Email: "email", CodeHash: "codehash"},
			err:   "expires is blank",
		},
		"returns nil for a valid invitation": {
			input: invitation.Invitation{Email: "email", CodeHash: "codehash", Expires: time.Now()},
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
