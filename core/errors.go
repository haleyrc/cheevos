package core

import (
	"fmt"
	"net/http"
)

type Model interface {
	Name() string
}

func NewValidationError(model Model) *ValidationError {
	return &ValidationError{
		Model:  model,
		Fields: map[string]string{},
	}
}

type ValidationError struct {
	Model  Model
	Fields map[string]string
}

func (ve *ValidationError) Add(fieldName, msg string) *ValidationError {
	ve.Fields[fieldName] = msg
	return ve
}

func (ve *ValidationError) Code() int { return http.StatusUnprocessableEntity }

func (ve *ValidationError) Error() error {
	if len(ve.Fields) > 0 {
		return validationError{ValidationError: ve}
	}
	return nil
}

func (ve *ValidationError) Message() string {
	return fmt.Sprintf("%s is invalid.", ve.Model.Name())
}

type validationError struct {
	*ValidationError
}

func (ve validationError) Error() string {
	return fmt.Sprintf("validation failed: %s is invalid", ve.Model.Name())
}
