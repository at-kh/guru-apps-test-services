package delivery

import "github.com/gofiber/fiber/v2"

type (
	// HealthHTTPHandler - describes an interface for work with services health over HTTP.
	HealthHTTPHandler interface {
		// Health - handler for getting meta information endpoint.
		Health(ctx *fiber.Ctx) error
	}

	// ProductsHTTPHandler - describes an interface for work with products over HTTP.
	ProductsHTTPHandler interface {
		// GetAll - handler for getting all products endpoint.
		GetAll(ctx *fiber.Ctx) error
		// Create - handler for creating product endpoint.
		Create(ctx *fiber.Ctx) error
		// Delete - handler for deleting product endpoint.
		Delete(ctx *fiber.Ctx) error
	}
)
