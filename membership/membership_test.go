package membership_test

import (
	"strings"
	"testing"

	"github.com/haleyrc/cheevos/lib/time"
	"github.com/haleyrc/cheevos/membership"
)

func TestValidatingAMembership(t *testing.T) {
	testcases := map[string]struct {
		input membership.Membership
		err   string
	}{
		"returns an error for a missing organization id": {
			input: membership.Membership{OrganizationID: "", UserID: "userid", Joined: time.Now()},
			err:   "organization id is blank",
		},
		"returns an error for a missing user id": {
			input: membership.Membership{OrganizationID: "orgid", UserID: "", Joined: time.Now()},
			err:   "user id is blank",
		},
		"returns an error for a missing joined time": {
			input: membership.Membership{OrganizationID: "orgid", UserID: "userid", Joined: time.Time{}},
			err:   "joined is blank",
		},
		"returns nil for a valid membership": {
			input: membership.Membership{OrganizationID: "orgid", UserID: "userid", Joined: time.Now()},
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
