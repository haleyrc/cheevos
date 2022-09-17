package membership_test

import (
	"testing"

	"github.com/haleyrc/cheevos/internal/testutil"
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
		testutil.RunValidationTests(t, name, &tc.input, tc.err)
	}
}
