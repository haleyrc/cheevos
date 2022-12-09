package domain

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

func (ae *AuthorizationError) Error() string { return fmt.Sprintf("not authorized: %v", ae.cause) }

func (ae *AuthorizationError) Message() string {
	msg := ae.msg
	if msg == "" {
		msg = "You are not permitted to perform that action."
	}
	return msg
}

func NewBadRequestError(err error) *BadRequestError {
	return &BadRequestError{cause: err}
}

type BadRequestError struct {
	cause error
}

func (bre *BadRequestError) Code() int { return http.StatusBadRequest }

func (bre *BadRequestError) Error() string { return fmt.Sprintf("bad request: %v", bre.cause) }

func (bre *BadRequestError) Message() string { return "There was a problem with your request." }

func NewRawError(code int, msg string) *RawError {
	return &RawError{
		code: code,
		msg:  msg,
	}
}

type RawError struct {
	code int
	msg  string
}

func (re *RawError) Code() int { return re.code }

func (re *RawError) Error() string { return fmt.Sprintf("error: %s", re.msg) }

func (re *RawError) Message() string { return re.msg }

type FieldError struct {
	Field string
	Msg   string
}

type ValidationError struct {
	Model       string
	FieldErrors []FieldError
}

func ValidationErrorFromError(err error) (ValidationError, bool) {
	ve, ok := err.(ValidationError)
	if !ok {
		return ValidationError{}, false
	}
	return ve, true
}

func NewValidationError(model string, fieldErrors []FieldError) ValidationError {
	return ValidationError{
		Model:       model,
		FieldErrors: fieldErrors,
	}
}

func (ve ValidationError) Code() int { return http.StatusUnprocessableEntity }

func (ve ValidationError) Error() string {
	return fmt.Sprintf("validation failed: %s is invalid", ve.Model)
}

func (ve ValidationError) FindFieldError(name string) (FieldError, bool) {
	for _, fe := range ve.FieldErrors {
		if fe.Field == name {
			return fe, true
		}
	}
	return FieldError{}, false
}

func (ve ValidationError) Message() string { return fmt.Sprintf("%s is invalid.", ve.Model) }
