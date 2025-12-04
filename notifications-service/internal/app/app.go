package app

import (
	"context"
	"time"

	"github.com/at-kh/guru-apps-test-services/notifications-service/internal/api/delivery"
	"github.com/at-kh/guru-apps-test-services/notifications-service/internal/api/domain/health"
	"github.com/at-kh/guru-apps-test-services/notifications-service/internal/api/services"
	"github.com/at-kh/guru-apps-test-services/notifications-service/internal/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type (
	// Meta defines meta for application.
	Meta struct {
		ConfigPath string
		Info       health.Info
	}

	// App defines main application struct.
	App struct {
		// meta information about application.
		meta Meta

		// Tech dependencies.
		cfg       *config.Config
		logger    *zap.Logger
		sqsClient *sqs.Client

		// Repository dependencies.

		// Services dependencies.
		notificationsService services.Notifications

		// Delivery dependencies.
		healthHTTPHandler  delivery.HealthHTTPHandler
		sqsConsumerHandler delivery.SQSConsumerHandler
	}

	worker func(ctx context.Context, a *App)
)

// New - app constructor without init for components.
func New(meta Meta) *App { return &App{meta: meta} }

// Run – initialize configuration and application dependencies, run workers.
func (a *App) Run(ctx context.Context) {
	decimal.MarshalJSONWithoutQuotes = true
	time.Local = time.UTC

	// Initialize configuration and tech dependencies
	a.initConfig()
	a.initLogger()
	a.initMessageBroker(ctx)

	// Layers registration
	a.registerServices()
	a.registerHTTPHandlers()
	a.registerBrokerHandlers()

	// Run workers
	a.runWorkers(ctx)
}
