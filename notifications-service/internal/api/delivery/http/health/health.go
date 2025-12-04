package health

import (
	"github.com/at-kh/guru-apps-test-services/notifications-service/internal/api/delivery"
	"github.com/at-kh/guru-apps-test-services/notifications-service/internal/api/domain/health"
	"github.com/gofiber/fiber/v2"
)

var _ delivery.HealthHTTPHandler = &Handler{}

type (
	// Handler defines a Handler for HTTP requests for checking status.
	Handler struct {
		infoResponse infoResponse
	}
)

// NewHandler defines a handler constructor.
func NewHandler(info health.Info) *Handler {
	return &Handler{
		infoResponse: infoResponse{
			Name:    info.Name,
			Commit:  info.BuildCommit,
			Date:    info.BuildDate,
			Version: info.BuildVersion,
		},
	}
}

// Health - handler for getting meta info endpoint.
func (h Handler) Health(ctx *fiber.Ctx) error { return ctx.JSON(h.infoResponse) }
