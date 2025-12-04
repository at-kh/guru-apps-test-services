package app

import (
	"github.com/at-kh/guru-apps-test-services/notifications-service/internal/api/delivery/http/health"
)

// registerHTTPHandlers initializes the http handlers.
func (a *App) registerHTTPHandlers() {
	a.healthHTTPHandler = health.NewHandler(a.meta.Info)
}
