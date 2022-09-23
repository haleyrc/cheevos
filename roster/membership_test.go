package roster_test

import (
	"testing"

	"github.com/haleyrc/cheevos/internal/testutil"
	"github.com/haleyrc/cheevos/lib/time"
	"github.com/haleyrc/cheevos/roster"
)

func TestValidatingAMembership(t *testing.T) {
	testcases := map[string]struct {
		input roster.Membership
		err   string
	}{
		"returns an error for a missing organization id": {
			input: roster.Membership{OrganizationID: "", UserID: "userid", Joined: time.Now()},
			err:   "organization id is blank",
		},
		"returns an error for a missing user id": {
			input: roster.Membership{OrganizationID: "orgid", UserID: "", Joined: time.Now()},
			err:   "user id is blank",
		},
		"returns an error for a missing joined time": {
			input: roster.Membership{OrganizationID: "orgid", UserID: "userid", Joined: time.Time{}},
			err:   "joined is blank",
		},
		"returns nil for a valid membership": {
			input: roster.Membership{OrganizationID: "orgid", UserID: "userid", Joined: time.Now()},
			err:   "",
		},
	}
	for name, tc := range testcases {
		testutil.RunValidationTests(t, name, &tc.input, tc.err)
	}
}
