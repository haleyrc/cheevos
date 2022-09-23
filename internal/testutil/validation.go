package testutil

import (
	"strings"
	"testing"
)

type Validator interface {
	Validate() error
}

func RunValidationTests(t *testing.T, name string, input Validator, err string) {
	t.Helper()
	t.Run(name, func(t *testing.T) {
		t.Helper()
		got := input.Validate()
		if err == "" {
			if got != nil {
				t.Error("unexpected error:", got)
			}
			return
		}

		if got == nil {
			t.Errorf("expected error, but got nil")
			return
		}

		compareError(t, err, got)
	})
}

func compareError(t *testing.T, want string, got error) {
	t.Helper()
	g := got.Error()
	if !strings.Contains(g, want) {
		t.Errorf("error %q does not include %q", g, want)
	}
}
