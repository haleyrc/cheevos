package core

import (
	"fmt"
	"net/http"
)

type Model interface {
	Model() string
}

func NewAuthorizationError(cause error, msg string) *AuthorizationError {
	return &AuthorizationError{
		cause: cause,
		msg:   msg,
	}
}

type AuthorizationError struct {
	cause error
	msg   string
}

func (ae *AuthorizationError) Code() int { return http.StatusForbidden }

func (ae *AuthorizationError) Error() string {
	return fmt.Sprintf("authorization failed: %v", ae.cause)
}

func (ae *AuthorizationError) Message() string {
	msg := ae.msg
	if msg == "" {
		msg = "You are not permitted to perform that action."
	}
	return msg
}

type ValidationError struct {
	Model  Model
	Fields map[string]string
}

func ValidationErrorFromError(err error) (*ValidationError, bool) {
	ve, ok := err.(validationError)
	if !ok {
		return nil, false
	}
	return ve.ValidationError, true
}

func NewValidationError(model Model) *ValidationError {
	return &ValidationError{
		Model:  model,
		Fields: map[string]string{},
	}
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
	return fmt.Sprintf("%s is invalid.", ve.Model.Model())
}

type validationError struct {
	*ValidationError
}

func (ve validationError) Error() string {
	return fmt.Sprintf("validation failed: %s is invalid", ve.Model.Model())
}
