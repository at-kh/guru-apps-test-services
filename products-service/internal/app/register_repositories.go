package app

import (
	"github.com/at-kh/guru-apps-test-services/products-service/internal/api/repositories/products"
	"github.com/at-kh/guru-apps-test-services/products-service/internal/api/repositories/sqs_publisher"
)

// registerRepositories registers repositories.
func (a *App) registerRepositories() {
	a.productsRepository = products.NewRepository(a.db, a.logger)
	a.sqsPublisherRepository = sqs_publisher.NewRepository(a.sqsClient, a.cfg.Delivery.Broker.URL, a.logger)
}
