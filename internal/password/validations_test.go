package password_test

import (
	"testing"

	"github.com/haleyrc/cheevos/internal/password"
)

func TestValidateLength(t *testing.T) {
	testcases := map[string]struct {
		input     password.Password
		shouldErr bool
	}{
		"should return an error if the password is too short": {
			input: password.New("123"), shouldErr: true,
		},
		"should return an error if the password is too long": {
			input: password.New("123456789"), shouldErr: true,
		},
		"should return nil for a good password": {
			input: password.New("12345678"), shouldErr: false,
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			validator := password.NewValidator(
				password.ValidateLength(4, 8),
			)

			err := validator.Validate(tc.input)

			if tc.shouldErr && err == nil {
				t.Error("Expected validation to fail, but it didn't.")
				return
			}

			if !tc.shouldErr && err != nil {
				t.Errorf("Expected validation to succeed, but got error %v.", err)
				return
			}
		})
	}
}
