package app

import (
	"github.com/at-kh/guru-apps-test-services/products-service/internal/api/services/products"
)

// registerServices register services in-app struct.
func (a *App) registerServices() {
	a.productsService = products.NewService(a.db, a.productsRepository, a.sqsPublisherRepository, a.metrics, a.logger)
}
