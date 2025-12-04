package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Namespace defines namespace for metrics.
const Namespace = "products_service"

type (
	// Metrics defines metrics for application.
	Metrics struct {
		ProductCreatedCounter prometheus.Counter
		ProductDeletedCounter prometheus.Counter
	}
)
