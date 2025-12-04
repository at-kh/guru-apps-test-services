package errs

import "net/http"

const (
	errInternal                = "internal_server_error"
	errNotImplemented          = "not_implemented"
	errBadGateway              = "bad_gateway"
	errServiceUnavailable      = "service_unavailable"
	errGatewayTimeout          = "gateway_timeout"
	errHTTPVersionNotSupported = "http_version_not_supported"
)

// HTTP server errors
type (
	// Internal - describes situation for internal server error.
	// Status code: 500
	Internal struct {
		Cause string `json:"cause"`
	}

	// NotImplemented - describes situation for not implemented.
	// Status code: 501
	NotImplemented struct {
		Cause string `json:"cause"`
	}

	// BadGateway - server error response code indicates that the server,
	// while acting as a gateway or proxy, received an invalid response from the upstream server.
	// Status code: 502
	BadGateway struct {
		Cause string `json:"cause"`
	}

	// ServiceUnavailable - describes situation for services unavailable.
	// Status code: 503
	ServiceUnavailable struct {
		Cause string `json:"cause"`
	}

	// GatewayTimeout - describes situation for gateway timeout.
	// Status code: 504
	GatewayTimeout struct {
		Cause string `json:"cause"`
	}

	// HTTPVersionNotSupported - describes situation for HTTP version not supported.
	// Status code: 505
	HTTPVersionNotSupported struct {
		Cause string `json:"cause"`
	}
)

// NewInternal - error for internal server error. Status code: 500
func NewInternal(cause string) Internal {
	return Internal{Cause: cause}
}

// Error implements error interface.
func (e Internal) Error() string {
	if e.Cause != "" {
		return format(errInternal, e.Cause)
	}

	return errInternal
}

// StatusCode implements error interface.
func (e Internal) StatusCode() int { return http.StatusInternalServerError }

// NewNotImplemented - error for not implemented. Status code: 501
func NewNotImplemented(cause string) NotImplemented {
	return NotImplemented{Cause: cause}
}

// Error implements error interface.
func (e NotImplemented) Error() string {
	if e.Cause != "" {
		return format(errNotImplemented, e.Cause)
	}

	return errNotImplemented
}

// StatusCode implements error interface.
func (e NotImplemented) StatusCode() int { return http.StatusNotImplemented }

// NewBadGateway - error for bad gateway. Status code: 502
func NewBadGateway(cause string) BadGateway {
	return BadGateway{Cause: cause}
}

// Error implements error interface.
func (e BadGateway) Error() string {
	if e.Cause != "" {
		return format(errBadGateway, e.Cause)
	}

	return errBadGateway
}

// StatusCode implements error interface.
func (e BadGateway) StatusCode() int { return http.StatusBadGateway }

// NewServiceUnavailable - error for services unavailable. Status code: 503
func NewServiceUnavailable(cause string) ServiceUnavailable {
	return ServiceUnavailable{Cause: cause}
}

// Error implements error interface.
func (e ServiceUnavailable) Error() string {
	if e.Cause != "" {
		return format(errServiceUnavailable, e.Cause)
	}

	return errServiceUnavailable
}

// StatusCode implements error interface.
func (e ServiceUnavailable) StatusCode() int { return http.StatusServiceUnavailable }

// NewGatewayTimeout - error for gateway timeout. Status code: 504
func NewGatewayTimeout(cause string) GatewayTimeout {
	return GatewayTimeout{Cause: cause}
}

// Error implements error interface.
func (e GatewayTimeout) Error() string {
	if e.Cause != "" {
		return format(errGatewayTimeout, e.Cause)
	}

	return errGatewayTimeout
}

// StatusCode implements error interface.
func (e GatewayTimeout) StatusCode() int { return http.StatusGatewayTimeout }

// NewHTTPVersionNotSupported - error for HTTP version not supported. Status code: 505
func NewHTTPVersionNotSupported(cause string) HTTPVersionNotSupported {
	return HTTPVersionNotSupported{Cause: cause}
}

// Error implements error interface.
func (e HTTPVersionNotSupported) Error() string {
	if e.Cause != "" {
		return format(errHTTPVersionNotSupported, e.Cause)
	}

	return errHTTPVersionNotSupported
}

// StatusCode implements error interface.
func (e HTTPVersionNotSupported) StatusCode() int { return http.StatusHTTPVersionNotSupported }
