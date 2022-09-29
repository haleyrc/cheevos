package core_test

import (
	"testing"

	"github.com/haleyrc/cheevos/core"
	"github.com/haleyrc/cheevos/lib/web"
)

func TestCoreErrorsAreCoded(t *testing.T) {
	testcases := map[string]interface{}{
		"authorization error": &core.AuthorizationError{},
		"raw error":           &core.RawError{},
		"validation error":    core.NewValidationError(testModel("Test")).Add("field", "msg").Error(),
	}
	for name, tc := range testcases {
		if _, ok := tc.(web.Coded); !ok {
			t.Errorf("Expected %s to be coded, but it isn't.", name)
		}
	}
}

func TestCoreErrorsAreMessaged(t *testing.T) {
	testcases := map[string]interface{}{
		"authorization error": &core.AuthorizationError{},
		"raw error":           &core.RawError{},
		"validation error":    core.NewValidationError(testModel("Test")).Add("field", "msg").Error(),
	}
	for name, tc := range testcases {
		if _, ok := tc.(web.Messaged); !ok {
			t.Errorf("Expected %s to be messaged, but it isn't.", name)
		}
	}
}

func TestValidationErrorAddsFieldErrors(t *testing.T) {
	fieldErrors := map[string]string{
		"MyField1": "My Field 1 shouldn't be blank",
		"MyField2": "My Field 2 shouldn't be blank",
	}

	ve := core.NewValidationError(testModel("Test"))

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
	var i interface{} = &core.ValidationError{}
	if err, ok := i.(error); ok {
		t.Errorf("Expected raw validation error to not be an error, but got %v.", err)
	}
}

func TestValidationErrorReturnsNilWithNoErrors(t *testing.T) {
	err := core.NewValidationError(testModel("Test")).Error()
	if err != nil {
		t.Errorf("Expected error to be nil, but got %v.", err)
	}
}

func TestValidationErrorReturnsAnErrorWithFieldErrors(t *testing.T) {
	ve := core.NewValidationError(testModel("Test"))
	ve.Add("MyField", "My Field shouldn't be blank")
	err := ve.Error()
	if err == nil {
		t.Errorf("Expected to get an error, but got nil.")
	}
}

type testModel string

func (tm testModel) Model() string { return string(tm) }
