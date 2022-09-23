package cheevos_test

import (
	"testing"

	"github.com/haleyrc/cheevos/cheevos"
	"github.com/haleyrc/cheevos/internal/testutil"
)

func TestNormalizingACheevoNormalizesName(t *testing.T) {
	subject := cheevos.Cheevo{Name: testutil.UnsafeString}
	subject.Normalize()
	if subject.Name != testutil.SafeString {
		t.Errorf("Expected cheevo name to be normalized, but it wasn't.")
	}
}

func TestNormalizingACheevoNormalizesDescription(t *testing.T) {
	subject := cheevos.Cheevo{Description: testutil.UnsafeString}
	subject.Normalize()
	if subject.Description != testutil.SafeString {
		t.Errorf("Expected cheevo description to be normalized, but it wasn't.")
	}
}

func TestValidatingACheevo(t *testing.T) {
	testcases := map[string]struct {
		input cheevos.Cheevo
		err   string
	}{
		"returns an error for a missing id": {
			input: cheevos.Cheevo{ID: "", Name: "name", Description: "description"},
			err:   "id is blank",
		},
		"returns an error for a missing name": {
			input: cheevos.Cheevo{ID: "id", Name: "", Description: "description"},
			err:   "name is blank",
		},
		"returns an error for a missing description": {
			input: cheevos.Cheevo{ID: "id", Name: "name", Description: ""},
			err:   "description is blank",
		},
		"returns nil for a valid cheevo": {
			input: cheevos.Cheevo{ID: "id", Name: "name", Description: "description"},
			err:   "",
		},
	}
	for name, tc := range testcases {
		testutil.RunValidationTests(t, name, &tc.input, tc.err)
	}
}
