package domain_test

import (
	"testing"

	"github.com/haleyrc/pkg/errors"

	"github.com/haleyrc/cheevos/domain"
	"github.com/haleyrc/cheevos/internal/assert"
)

func TestAllErrorsAreCoded(t *testing.T) {
	testcases := map[string]interface{}{
		"authorization error": &domain.AuthorizationError{},
		"bad request error":   &domain.BadRequestError{},
		"raw error":           &domain.RawError{},
		"validation error":    domain.NewValidationError(testModel("Test")).Add("field", "msg").Error(),
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
		"validation error":    domain.NewValidationError(testModel("Test")).Add("field", "msg").Error(),
	}
	for name, tc := range testcases {
		if _, ok := tc.(errors.Messager); !ok {
			t.Errorf("Expected %s to be messaged, but it isn't.", name)
		}
	}
}

func TestValidationErrorAddsFieldErrors(t *testing.T) {
	fieldErrors := map[string]string{
		"MyField1": "My Field 1 shouldn't be blank",
		"MyField2": "My Field 2 shouldn't be blank",
	}

	ve := domain.NewValidationError(testModel("Test"))

	for name, msg := range fieldErrors {
		ve.Add(name, msg)
	}

	for name, want := range fieldErrors {
		got, ok := ve.Fields[name]
		if !ok {
			t.Errorf("Expected validation error to contain field error for %q, but it didn't.", name)
		} else {
			if got != want {
				t.Errorf("Expected field error message to be %q, but got %q.", want, got)
			}
		}
	}
}

func TestValidationErrorIsntAnError(t *testing.T) {
	var i interface{} = &domain.ValidationError{}
	if err, ok := i.(error); ok {
		t.Errorf("Expected raw validation error to not be an error, but got %v.", err)
	}
}

func TestValidationErrorReturnsNilWithNoErrors(t *testing.T) {
	err := domain.NewValidationError(testModel("Test")).Error()
	assert.Error(t, err).IsNil()
}

func TestValidationErrorReturnsAnErrorWithFieldErrors(t *testing.T) {
	ve := domain.NewValidationError(testModel("Test"))
	ve.Add("MyField", "My Field shouldn't be blank")
	assert.Error(t, ve.Error()).IsNotNil()
}

type testModel string

func (tm testModel) Model() string { return string(tm) }
