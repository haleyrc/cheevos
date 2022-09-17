package cheevo_test

import (
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
		testutil.RunValidationTests(t, name, &tc.input, tc.err)
	}
}
