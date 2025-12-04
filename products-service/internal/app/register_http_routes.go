package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

// registerHTTPRoutes registers http routes.
func (a *App) registerHTTPRoutes(app *fiber.App) {
	r := app.Group("/products-api/v1")
	r.Get("/health", a.healthHTTPHandler.Health)

	products := r.Group("/products")
	products.Get("", a.productsHTTPHandler.GetAll)
	products.Post("", a.productsHTTPHandler.Create)
	products.Delete("/:id", a.productsHTTPHandler.Delete)

	r.Get("/metrics", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler())(c.Context())
		return c.SendStatus(fiber.StatusOK)
	})
}
