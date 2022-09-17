package cheevo_test

import (
	"strings"
	"testing"

	"github.com/haleyrc/cheevos/cheevo"
	"github.com/haleyrc/cheevos/internal/testutil"
)

func TestNormalizingACheevoNormalizesName(t *testing.T) {
	subject := cheevo.Cheevo{Name: testutil.UnsafeString}
	subject.Normalize()
	if subject.Name != testutil.SafeString {
		t.Errorf("Expected cheevo name to be normalized, but it wasn't.")
	}
}

func TestNormalizingACheevoNormalizesDescription(t *testing.T) {
	subject := cheevo.Cheevo{Description: testutil.UnsafeString}
	subject.Normalize()
	if subject.Description != testutil.SafeString {
		t.Errorf("Expected cheevo description to be normalized, but it wasn't.")
	}
}

func TestValidatingACheevo(t *testing.T) {
	testcases := map[string]struct {
		input cheevo.Cheevo
		err   string
	}{
		"returns an error for a missing id": {
			input: cheevo.Cheevo{ID: "", Name: "name", Description: "description"},
			err:   "id is blank",
		},
		"returns an error for a missing name": {
			input: cheevo.Cheevo{ID: "id", Name: "", Description: "description"},
			err:   "name is blank",
		},
		"returns an error for a missing description": {
			input: cheevo.Cheevo{ID: "id", Name: "name", Description: ""},
			err:   "description is blank",
		},
		"returns nil for a valid cheevo": {
			input: cheevo.Cheevo{ID: "id", Name: "name", Description: "description"},
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
