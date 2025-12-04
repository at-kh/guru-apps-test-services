package app

import (
	"github.com/at-kh/guru-apps-test-services/products-service/internal/api/domain/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// initMetrics - initialize metrics for application.
func (a *App) initMetrics() {
	a.metrics = &metrics.Metrics{
		ProductCreatedCounter: promauto.NewCounter(
			prometheus.CounterOpts{
				Namespace: metrics.Namespace,
				Name:      "created_products_cnt",
				Help:      "Total number of products created",
			},
		),
		ProductDeletedCounter: promauto.NewCounter(
			prometheus.CounterOpts{
				Namespace: metrics.Namespace,
				Name:      "deleted_products_cnt",
				Help:      "Total number of products deleted",
			},
		),
	}
}
