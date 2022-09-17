package award_test

import (
	"testing"

	"github.com/haleyrc/cheevos/award"
	"github.com/haleyrc/cheevos/internal/testutil"
	"github.com/haleyrc/cheevos/lib/time"
)

func TestValidatingAAward(t *testing.T) {
	testcases := map[string]struct {
		input award.Award
		err   string
	}{
		"returns an error for a missing cheevo id": {
			input: award.Award{CheevoID: "", UserID: "userid", Awarded: time.Now()},
			err:   "cheevo id is blank",
		},
		"returns an error for a missing user id": {
			input: award.Award{CheevoID: "orgid", UserID: "", Awarded: time.Now()},
			err:   "user id is blank",
		},
		"returns an error for a missing awarded time": {
			input: award.Award{CheevoID: "orgid", UserID: "userid", Awarded: time.Time{}},
			err:   "awarded is blank",
		},
		"returns nil for a valid membership": {
			input: award.Award{CheevoID: "orgid", UserID: "userid", Awarded: time.Now()},
			err:   "",
		},
	}
	for name, tc := range testcases {
		testutil.RunValidationTests(t, name, &tc.input, tc.err)
	}
}
