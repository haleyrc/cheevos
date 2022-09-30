package core

import (
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"

	"github.com/pkg/errors"
)

type Coded interface {
	Code() int
}

type Messaged interface {
	Message() string
}

type Model interface {
	Model() string
}

func ErrorCode(err error) int {
	root := errors.Cause(err)
	return errorCode(root)
}

func errorCode(err error) int {
	if err, ok := err.(Coded); ok {
		return err.Code()
	}
	return http.StatusInternalServerError
}

func ErrorMessage(err error) string {
	root := errors.Cause(err)
	return errorMessage(root)
}

func errorMessage(err error) string {
	if err, ok := err.(Messaged); ok {
		return err.Message()
	}
	return "Something went wrong."
}

type wrappedError struct {
	file   string
	lineno int
	name   string
	cause  error
}

func (we wrappedError) Code() int { return ErrorCode(we.cause) }

func (we wrappedError) Error() string { return fmt.Sprintf("%s: %v", we.name, we.cause) }

func (we wrappedError) Message() string { return ErrorMessage(we.cause) }

func WrapError(err error) error {
	pc, file, lineno, ok := runtime.Caller(1)

	we := wrappedError{
		file:   filepath.Base(file),
		lineno: lineno,
		cause:  err,
	}

	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		we.name = filepath.Base(details.Name())
	}

	return we
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

func (ve *ValidationError) Message() string { return fmt.Sprintf("%s is invalid.", ve.Model.Model()) }

type validationError struct {
	*ValidationError
}

func (ve validationError) Error() string {
	return fmt.Sprintf("validation failed: %s is invalid", ve.Model.Model())
}
