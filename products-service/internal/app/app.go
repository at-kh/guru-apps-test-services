package app

import (
	"context"
	"embed"
	"time"

	"github.com/at-kh/guru-apps-test-services/products-service/internal/api/delivery"
	"github.com/at-kh/guru-apps-test-services/products-service/internal/api/domain/health"
	"github.com/at-kh/guru-apps-test-services/products-service/internal/api/domain/metrics"
	"github.com/at-kh/guru-apps-test-services/products-service/internal/api/repositories"
	"github.com/at-kh/guru-apps-test-services/products-service/internal/api/services"
	"github.com/at-kh/guru-apps-test-services/products-service/internal/config"
	"github.com/at-kh/guru-apps-test-services/products-service/pkg/responder"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/jmoiron/sqlx"
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

		// FS dependencies.
		dbMigrationsFS embed.FS

		// Tech dependencies.
		cfg       *config.Config
		logger    *zap.Logger
		db        *sqlx.DB
		responder responder.Responder // responder for http responses
		sqsClient *sqs.Client

		// Metrics dependencies.
		metrics *metrics.Metrics

		// Repository dependencies.
		productsRepository     repositories.ProductsRepository
		sqsPublisherRepository repositories.SQSPublisherRepository

		// Services dependencies.
		productsService services.ProductsService

		// Delivery dependencies.
		healthHTTPHandler   delivery.HealthHTTPHandler
		productsHTTPHandler delivery.ProductsHTTPHandler
	}

	worker func(ctx context.Context, a *App)
)

// New - app constructor without init for components.
func New(meta Meta) *App { return &App{meta: meta} }

// WithMigrationFS is a setup for database migration filesystem
func (a *App) WithMigrationFS(dbMigrationFS embed.FS) *App {
	a.dbMigrationsFS = dbMigrationFS
	return a
}

// Run – initialize configuration and application dependencies, run workers.
func (a *App) Run(ctx context.Context) {
	decimal.MarshalJSONWithoutQuotes = true
	time.Local = time.UTC

	// Initialize configuration and tech dependencies
	a.initConfig()
	a.initLogger()
	a.initMetrics()
	a.initResponder()
	a.initDatabase()
	a.initMessageBroker(ctx)

	// Layers registration
	a.registerRepositories()
	a.registerServices()
	a.registerHTTPHandlers()

	// Run workers
	a.runWorkers(ctx)
}
