package testutil

import (
	"strings"
	"testing"

	"github.com/haleyrc/cheevos"
)

type Validator interface{ Validate() error }

func RunValidationTests(t *testing.T, v Validator, msg string, fieldErrors map[string]string) {
	t.Helper()

	err := v.Validate()
	if err == nil {
		t.Fatal("Expected validate to return an error, but it didn't.")
	}

	if got := err.Error(); got != msg {
		t.Errorf("Expected error to be %q, but got %q.", msg, got)
	}

	ve, ok := cheevos.ValidationErrorFromError(err)
	if !ok {
		t.Fatalf("Expected error to be a validation error, but it was a(n) %T.", err)
	}

	for name, msg := range fieldErrors {
		got, ok := ve.Fields[name]
		if !ok {
			t.Fatalf("Expected field errors to include %s, but they didn't.", name)
		}
		if got != msg {
			t.Errorf("Expected field error to be %q, but got %q.", msg, got)
		}
	}
}

func CompareError(t *testing.T, want string, got error) bool {
	t.Helper()
	g := got.Error()
	if !strings.Contains(g, want) {
		t.Errorf("error %q does not include %q", g, want)
		return false
	}
	return true
}
