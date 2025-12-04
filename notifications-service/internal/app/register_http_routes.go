package app

import (
	"github.com/gofiber/fiber/v2"
)

// registerHTTPRoutes registers http routes.
func (a *App) registerHTTPRoutes(app *fiber.App) {
	r := app.Group("/notifications-api/v1")
	r.Get("/health", a.healthHTTPHandler.Health)
}
