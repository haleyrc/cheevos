package domain_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/haleyrc/pkg/errors"

	"github.com/haleyrc/cheevos/domain"
)

func TestStatusCodeReturnsTheCorrectCode(t *testing.T) {
	testcases := []struct {
		err  error
		want int
	}{
		{domain.AuthorizationError{}, http.StatusForbidden},
		{domain.BadRequestError{}, http.StatusBadRequest},
		{domain.InvitationExpiredError{}, http.StatusGone},
		{domain.ValidationError{}, http.StatusUnprocessableEntity},
		{fmt.Errorf("oops"), http.StatusInternalServerError},
	}
	for _, tc := range testcases {
		got := domain.StatusCode(tc.err)
		if got != tc.want {
			t.Errorf("Expected %T to have status code %d, but got %d.", tc.err, tc.want, got)
		}
	}
}

func TestAllErrorsAreMessaged(t *testing.T) {
	testcases := map[string]interface{}{
		"authorization error":      &domain.AuthorizationError{},
		"bad request error":        &domain.BadRequestError{},
		"invitation expired error": &domain.InvitationExpiredError{},
		"validation error":         domain.NewValidationError("Test", nil),
	}
	for name, tc := range testcases {
		if _, ok := tc.(errors.Messager); !ok {
			t.Errorf("Expected %s to be messaged, but it isn't.", name)
		}
	}
}
