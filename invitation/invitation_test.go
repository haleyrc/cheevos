package invitation_test

import (
	"testing"

	"github.com/haleyrc/cheevos/internal/testutil"
	"github.com/haleyrc/cheevos/invitation"
	"github.com/haleyrc/cheevos/lib/time"
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
			input: invitation.Invitation{Email: "", OrganizationID: "orgid", Expires: time.Now()},
			err:   "email is blank",
		},
		"returns an error for a missing organization id": {
			input: invitation.Invitation{Email: "email", OrganizationID: "", Expires: time.Now()},
			err:   "organization id is blank",
		},
		"returns an error for a missing expiration": {
			input: invitation.Invitation{Email: "email", OrganizationID: "orgid"},
			err:   "expires is blank",
		},
		"returns nil for a valid invitation": {
			input: invitation.Invitation{Email: "email", OrganizationID: "orgid", Expires: time.Now()},
			err:   "",
		},
	}
	for name, tc := range testcases {
		testutil.RunValidationTests(t, name, &tc.input, tc.err)
	}
}
