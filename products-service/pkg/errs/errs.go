package errs

import (
	"net/http"
	"strings"
)

const (
	errEmpty      = "empty"
	errValidation = "validation_error"

	divider   = " - "
	separator = ","
)

// HTTPError - interface for HTTP errors
type HTTPError interface {
	error
	StatusCode() int
}

// Custom errors
type (
	// Empty - an error that indicates that nothing was found by the specified search parameters.
	Empty struct {
		What string `json:"what"`
	}

	// FieldsValidation - describes situation for bad request with text.
	FieldsValidation struct {
		Errors []string `json:"errors"`
	}
)

// NewEmpty - error for empty.
func NewEmpty(what string) Empty { return Empty{What: what} }

// Error implements error interface.
func (e Empty) Error() string {
	if e.What != "" {
		return format(errEmpty, e.What)
	}

	return errEmpty
}

// StatusCode implements error interface.
func (e Empty) StatusCode() int { return http.StatusNotFound }

// NewFieldsValidation - error for fields validation.
func NewFieldsValidation(errors []string) FieldsValidation {
	return FieldsValidation{Errors: errors}
}

// Error implements error interface.
func (e FieldsValidation) Error() string {
	if len(e.Errors) == 0 {
		return errValidation
	}

	var sb strings.Builder
	sb.WriteString(errValidation)
	sb.WriteString(divider)
	sb.WriteString(strings.Join(e.Errors, separator))

	return sb.String()
}

// StatusCode implements error interface.
func (e FieldsValidation) StatusCode() int { return http.StatusBadRequest }

// format formats error
func format(messages ...string) string { return strings.Join(messages, divider) }
