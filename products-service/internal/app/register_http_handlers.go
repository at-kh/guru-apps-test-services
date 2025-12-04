package app

import (
	"github.com/at-kh/guru-apps-test-services/products-service/internal/api/delivery/http/health"
	"github.com/at-kh/guru-apps-test-services/products-service/internal/api/delivery/http/products"
)

// registerHTTPHandlers initializes the http handlers.
func (a *App) registerHTTPHandlers() {
	a.productsHTTPHandler = products.NewHandler(a.responder, a.productsService, a.logger)
	a.healthHTTPHandler = health.NewHandler(a.meta.Info)
}
