package domain_test

import (
	"testing"

	"github.com/haleyrc/pkg/errors"

	"github.com/haleyrc/cheevos/domain"
)

func TestAllErrorsAreCoded(t *testing.T) {
	testcases := map[string]interface{}{
		"authorization error": &domain.AuthorizationError{},
		"bad request error":   &domain.BadRequestError{},
		"raw error":           &domain.RawError{},
		"validation error":    domain.NewValidationError("Test", nil),
	}
	for name, tc := range testcases {
		if _, ok := tc.(errors.Coder); !ok {
			t.Errorf("Expected %s to be coded, but it isn't.", name)
		}
	}
}

func TestAllErrorsAreMessaged(t *testing.T) {
	testcases := map[string]interface{}{
		"authorization error": &domain.AuthorizationError{},
		"bad request error":   &domain.BadRequestError{},
		"raw error":           &domain.RawError{},
		"validation error":    domain.NewValidationError("Test", nil),
	}
	for name, tc := range testcases {
		if _, ok := tc.(errors.Messager); !ok {
			t.Errorf("Expected %s to be messaged, but it isn't.", name)
		}
	}
}
