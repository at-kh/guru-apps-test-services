package responder

import (
	"errors"
	"net/http"

	"github.com/at-kh/guru-apps-test-services/products-service/pkg/errs"
	"github.com/gofiber/fiber/v2"
)

type (
	// Error - struct for error response.
	Error struct {
		Error string `json:"error"`
	}

	// Responder - helper struct wraps request/response functionality for fiber http handlers.
	Responder struct{}
)

// New - constructor for Responder entity.
func New() Responder { return Responder{} }

// Respond - helper func for respond.
func (h Responder) Respond(ctx *fiber.Ctx, code int, payload any) error {
	ctx.Response().SetStatusCode(code)

	if err := ctx.JSON(payload); err != nil {
		return errs.Internal{Cause: "invalid json in output"}
	}

	return nil
}

// RespondError - helper func for response with error.
func (h Responder) RespondError(ctx *fiber.Ctx, code int, err Error) error {
	return h.Respond(ctx, code, err)
}

// RespondEmpty - helper func for empty response.
func (h Responder) RespondEmpty(ctx *fiber.Ctx, code int) error {
	ctx.Response().SetStatusCode(code)
	return nil
}

// HandleError - helper func for response with error.
func (h Responder) HandleError(ctx *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, fiber.ErrMethodNotAllowed):
		return h.Respond(ctx, http.StatusMethodNotAllowed, Error{Error: err.Error()})
	case errors.Is(err, fiber.ErrNotImplemented):
		return h.Respond(ctx, http.StatusNotImplemented, Error{Error: err.Error()})
	case errors.Is(err, fiber.ErrBadGateway):
		return h.Respond(ctx, http.StatusBadGateway, Error{Error: err.Error()})
	case errors.Is(err, fiber.ErrServiceUnavailable):
		return h.Respond(ctx, http.StatusServiceUnavailable, Error{Error: err.Error()})
	case errors.Is(err, fiber.ErrGatewayTimeout):
		return h.Respond(ctx, http.StatusGatewayTimeout, Error{Error: err.Error()})
	case errors.Is(err, fiber.ErrHTTPVersionNotSupported):
		return h.Respond(ctx, http.StatusHTTPVersionNotSupported, Error{Error: err.Error()})
	}

	var httpErr errs.HTTPError
	if errors.As(err, &httpErr) {
		if httpErr.StatusCode() == http.StatusNoContent {
			return h.RespondEmpty(ctx, http.StatusNoContent)
		}

		return h.Respond(ctx, httpErr.StatusCode(), Error{Error: httpErr.Error()})
	}

	return h.Respond(ctx, http.StatusInternalServerError, Error{Error: errs.Internal{Cause: err.Error()}.Error()})
}
