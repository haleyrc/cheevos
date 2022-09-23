package roster_test

import (
	"testing"

	"github.com/haleyrc/cheevos/internal/testutil"
	"github.com/haleyrc/cheevos/roster"
)

func TestNormalizingAnOrganizationNormalizesName(t *testing.T) {
	subject := roster.Organization{Name: testutil.UnsafeString}
	subject.Normalize()
	if subject.Name != testutil.SafeString {
		t.Errorf("Expected organization name to be normalized, but it wasn't.")
	}
}

func TestValidatingAnOrganization(t *testing.T) {
	testcases := map[string]struct {
		input roster.Organization
		err   string
	}{
		"returns an error for a missing id": {
			input: roster.Organization{ID: "", Name: "name", OwnerID: "owner"},
			err:   "id is blank",
		},
		"returns an error for a missing name": {
			input: roster.Organization{ID: "id", Name: "", OwnerID: "owner"},
			err:   "name is blank",
		},
		"returns an error for a missing owner id": {
			input: roster.Organization{ID: "id", Name: "name", OwnerID: ""},
			err:   "owner id is blank",
		},
		"returns nil for a valid organization": {
			input: roster.Organization{ID: "id", Name: "name", OwnerID: "owner"},
			err:   "",
		},
	}
	for name, tc := range testcases {
		testutil.RunValidationTests(t, name, &tc.input, tc.err)
	}
}
