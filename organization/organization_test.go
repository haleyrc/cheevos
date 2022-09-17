package organization_test

import (
	"strings"
	"testing"

	"github.com/haleyrc/cheevos/internal/testutil"
	"github.com/haleyrc/cheevos/organization"
)

func TestNormalizingAnOrganizationNormalizesName(t *testing.T) {
	subject := organization.Organization{Name: testutil.UnsafeString}
	subject.Normalize()
	if subject.Name != testutil.SafeString {
		t.Errorf("Expected organization name to be normalized, but it wasn't.")
	}
}

func TestValidatingAnOrganization(t *testing.T) {
	testcases := map[string]struct {
		input organization.Organization
		err   string
	}{
		"returns an error for a missing id": {
			input: organization.Organization{ID: "", Name: "name", OwnerID: "owner"},
			err:   "id is blank",
		},
		"returns an error for a missing name": {
			input: organization.Organization{ID: "id", Name: "", OwnerID: "owner"},
			err:   "name is blank",
		},
		"returns an error for a missing owner id": {
			input: organization.Organization{ID: "id", Name: "name", OwnerID: ""},
			err:   "owner id is blank",
		},
		"returns nil for a valid organization": {
			input: organization.Organization{ID: "id", Name: "name", OwnerID: "owner"},
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
