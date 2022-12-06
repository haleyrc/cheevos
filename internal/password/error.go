package password

import "strings"

type ValidationError struct {
	Errors []error
}

func (ve ValidationError) Error() string {
	errorStrings := make([]string, 0, len(ve.Errors))
	for _, err := range ve.Errors {
		errorStrings = append(errorStrings, err.Error())
	}
	return strings.Join(errorStrings, "; ")
}
