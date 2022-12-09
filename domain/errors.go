package domain

import (
	"fmt"
	"net/http"
)

func StatusCode(err error) int {
	switch err.(type) {
	case AuthorizationError:
		return http.StatusForbidden
	case BadRequestError:
		return http.StatusBadRequest
	case InvitationExpiredError:
		return http.StatusGone
	case ValidationError:
		return http.StatusUnprocessableEntity
	default:
		return http.StatusInternalServerError
	}
}

type Model interface {
	Model() string
}

func NewAuthorizationError(cause error, msg string) AuthorizationError {
	return AuthorizationError{
		cause: cause,
		msg:   msg,
	}
}

type AuthorizationError struct {
	cause error
	msg   string
}

func (ae AuthorizationError) Error() string { return fmt.Sprintf("not authorized: %v", ae.cause) }

func (ae AuthorizationError) Message() string {
	msg := ae.msg
	if msg == "" {
		msg = "You are not permitted to perform that action."
	}
	return msg
}

func NewBadRequestError(err error) BadRequestError {
	return BadRequestError{cause: err}
}

type BadRequestError struct {
	cause error
}

func (bre BadRequestError) Error() string { return fmt.Sprintf("bad request: %v", bre.cause) }

func (bre BadRequestError) Message() string { return "There was a problem with your request." }

type InvitationExpiredError struct{}

func (iee InvitationExpiredError) Error() string { return "invitation expired" }

func (ieee InvitationExpiredError) Message() string {
	return "Your invitation has expired. Please contact your organization administrator for a new invitation."
}

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
