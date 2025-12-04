package errs

import "net/http"

const (
	errUnauthorized         = "unauthorized"
	errPaymentRequired      = "payment_required"
	errForbidden            = "forbidden"
	errNotFound             = "not_found"
	errMethodNotAllowed     = "method_not_allowed"
	errNotAcceptable        = "not_acceptable"
	errConflict             = "conflict"
	errRequestTimeout       = "request_timeout"
	errGone                 = "gone"
	errPayloadTooLarge      = "payload_too_large"
	errUnsupportedMediaType = "unsupported_media_type"
	errTooManyRequests      = "too_many_requests"
)

// HTTP client errors
type (
	// BadRequest - describes situation for bad request with text.
	// Status code: 400
	BadRequest struct {
		Cause string `json:"cause"`
	}

	// Unauthorized - the error determines the situation when the user is not authorized for the specified operation.
	// Status code: 401
	Unauthorized struct {
		Cause string `json:"cause"`
	}

	// PaymentRequired - describes situation for payment required.
	// Status code: 402
	PaymentRequired struct {
		Cause string `json:"cause"`
	}

	// Forbidden - describes situation for forbidden access to operation.
	// Status code: 403
	Forbidden struct {
		Cause string `json:"cause"`
	}

	// NotFound - describes situation for any item not found.
	// Status code: 404
	NotFound struct {
		What string `json:"what"`
	}

	// MethodNotAllowed - describes situation for any method not allowed.
	// Status code: 405
	MethodNotAllowed struct {
		Cause string `json:"cause"`
	}

	// NotAcceptable - describes situation for not acceptable.
	// Status code: 406
	NotAcceptable struct {
		Cause string `json:"cause"`
	}

	// RequestTimeout - describes situation for request timeout.
	// Status code: 408
	RequestTimeout struct {
		Cause string `json:"cause"`
	}

	// Conflict - describes situation for conflict in logic with current application state.
	// Status code: 409
	Conflict struct {
		What string `json:"what"`
	}

	// Gone - describes situation for item already deleted.
	// Status code: 410
	Gone struct {
		What string `json:"what"`
	}

	// PayloadTooLarge - describes situation for request entity too large.
	// Status code: 413
	PayloadTooLarge struct {
		Cause string `json:"cause"`
	}

	// UnsupportedMediaType - describes situation for unsupported media type.
	// Status code: 415
	UnsupportedMediaType struct {
		Cause string `json:"cause"`
	}

	// TooManyRequests - describes situation for too many requests.
	// Status code: 429
	TooManyRequests struct {
		Cause string `json:"cause"`
	}
)

// NewBadRequest - error for bad request. Status code: 400
func NewBadRequest(cause string) BadRequest {
	return BadRequest{Cause: cause}
}

// Error implements error interface.
func (e BadRequest) Error() string {
	return e.Cause
}

// StatusCode implements error interface.
func (e BadRequest) StatusCode() int { return http.StatusBadRequest }

// NewUnauthorized - error for unauthorized. Status code: 401
func NewUnauthorized(cause string) Unauthorized {
	return Unauthorized{Cause: cause}
}

// Error implements error interface.
func (e Unauthorized) Error() string {
	if e.Cause != "" {
		return format(errUnauthorized, e.Cause)
	}

	return errUnauthorized
}

// StatusCode implements error interface.
func (e Unauthorized) StatusCode() int { return http.StatusUnauthorized }

// NewPaymentRequired - error for payment required. Status code: 402
func NewPaymentRequired(cause string) PaymentRequired {
	return PaymentRequired{Cause: cause}
}

// Error implements error interface.
func (e PaymentRequired) Error() string {
	if e.Cause != "" {
		return format(errPaymentRequired, e.Cause)
	}

	return errPaymentRequired
}

// StatusCode implements error interface.
func (e PaymentRequired) StatusCode() int { return http.StatusPaymentRequired }

// NewForbidden - error for forbidden. Status code: 403
func NewForbidden(cause string) Forbidden {
	return Forbidden{Cause: cause}
}

// Error implements error interface.
func (e Forbidden) Error() string {
	if e.Cause != "" {
		return format(errForbidden, e.Cause)
	}

	return errForbidden
}

// StatusCode implements error interface.
func (e Forbidden) StatusCode() int { return http.StatusForbidden }

// NewNotFound - error for not found. Status code: 404
func NewNotFound(what string) NotFound {
	return NotFound{What: what}
}

// Error implements error interface.
func (e NotFound) Error() string {
	if e.What != "" {
		return format(errNotFound, e.What)
	}

	return errNotFound
}

// StatusCode implements error interface.
func (e NotFound) StatusCode() int { return http.StatusNotFound }

// NewMethodNotAllowed - error for method not allowed. Status code: 405
func NewMethodNotAllowed(cause string) MethodNotAllowed {
	return MethodNotAllowed{Cause: cause}
}

// Error implements error interface.
func (e MethodNotAllowed) Error() string {
	if e.Cause != "" {
		return format(errMethodNotAllowed, e.Cause)
	}

	return errMethodNotAllowed
}

// StatusCode implements error interface.
func (e MethodNotAllowed) StatusCode() int { return http.StatusMethodNotAllowed }

// NewNotAcceptable - error for not acceptable. Status code: 406
func NewNotAcceptable(cause string) NotAcceptable {
	return NotAcceptable{Cause: cause}
}

// Error implements error interface.
func (e NotAcceptable) Error() string {
	if e.Cause != "" {
		return format(errNotAcceptable, e.Cause)
	}

	return errNotAcceptable
}

// StatusCode implements error interface.
func (e NotAcceptable) StatusCode() int { return http.StatusNotAcceptable }

// NewRequestTimeout - error for request timeout. Status code: 408
func NewRequestTimeout(cause string) RequestTimeout {
	return RequestTimeout{Cause: cause}
}

// Error implements error interface.
func (e RequestTimeout) Error() string {
	if e.Cause != "" {
		return format(errRequestTimeout, e.Cause)
	}

	return errRequestTimeout
}

// StatusCode implements error interface.
func (e RequestTimeout) StatusCode() int { return http.StatusRequestTimeout }

// NewConflict - error for conflict. Status code: 409
func NewConflict(what string) Conflict {
	return Conflict{What: what}
}

// Error implements error interface.
func (e Conflict) Error() string {
	if e.What != "" {
		return format(errConflict, e.What)
	}

	return errConflict
}

// StatusCode implements error interface.
func (e Conflict) StatusCode() int { return http.StatusConflict }

// NewGone - error for gone. Status code: 410
func NewGone(what string) Gone {
	return Gone{What: what}
}

// Error implements error interface.
func (e Gone) Error() string {
	if e.What != "" {
		return e.What
	}

	return errGone
}

// StatusCode implements error interface.
func (e Gone) StatusCode() int { return http.StatusGone }

// NewPayloadTooLarge - error for payload too large. Status code: 413
func NewPayloadTooLarge(cause string) PayloadTooLarge {
	return PayloadTooLarge{Cause: cause}
}

// Error implements error interface.
func (e PayloadTooLarge) Error() string {
	if e.Cause != "" {
		return format(errPayloadTooLarge, e.Cause)
	}

	return errPayloadTooLarge
}

// StatusCode implements error interface.
func (e PayloadTooLarge) StatusCode() int { return http.StatusRequestEntityTooLarge }

// NewUnsupportedMediaType - error for unsupported media type. Status code: 415
func NewUnsupportedMediaType(cause string) UnsupportedMediaType {
	return UnsupportedMediaType{Cause: cause}
}

// Error implements error interface.
func (e UnsupportedMediaType) Error() string {
	if e.Cause != "" {
		return format(errUnsupportedMediaType, e.Cause)
	}

	return errUnsupportedMediaType
}

// StatusCode implements error interface.
func (e UnsupportedMediaType) StatusCode() int { return http.StatusUnsupportedMediaType }

// NewTooManyRequests - error for too many requests. Status code: 429
func NewTooManyRequests(cause string) TooManyRequests {
	return TooManyRequests{Cause: cause}
}

// Error implements error interface.
func (e TooManyRequests) Error() string {
	if e.Cause != "" {
		return format(errTooManyRequests, e.Cause)
	}

	return errTooManyRequests
}

// StatusCode implements error interface.
func (e TooManyRequests) StatusCode() int { return http.StatusTooManyRequests }
