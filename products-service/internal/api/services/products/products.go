package products

import (
	"context"
	"database/sql"

	"github.com/at-kh/guru-apps-test-services/products-service/internal/api/domain/metrics"
	"github.com/at-kh/guru-apps-test-services/products-service/internal/api/domain/products"
	"github.com/at-kh/guru-apps-test-services/products-service/internal/api/repositories"
	repoproducts "github.com/at-kh/guru-apps-test-services/products-service/internal/api/repositories/products"
	"github.com/at-kh/guru-apps-test-services/products-service/internal/api/services"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var _ services.ProductsService = &Service{}

// Service - defines services struct.
type Service struct {
	db                     *sqlx.DB
	productsRepository     repositories.ProductsRepository
	sqsPublisherRepository repositories.SQSPublisherRepository
	metrics                *metrics.Metrics
	logger                 *zap.Logger
}

// NewService constructor.
func NewService(
	db *sqlx.DB,
	productsRepository repositories.ProductsRepository,
	sqsPublisherRepository repositories.SQSPublisherRepository,
	metrics *metrics.Metrics,
	logger *zap.Logger,
) *Service {
	return &Service{
		db:                     db,
		productsRepository:     productsRepository,
		sqsPublisherRepository: sqsPublisherRepository,
		metrics:                metrics,
		logger:                 logger.With(zap.String("services", "products")),
	}
}

// Create creates a new product.
func (s Service) Create(ctx context.Context, p products.Product) (product products.Product, err error) {
	if ctx.Err() != nil {
		return products.Product{}, ctx.Err()
	}

	tx, err := s.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		s.logger.Error("failed to begin transaction", zap.Error(err))
		return products.Product{}, err
	}
	defer func() {
		if pp := recover(); pp != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				s.logger.Error("failed to rollback transaction after panic", zap.Error(rollbackErr))
			}
			panic(pp)
		}
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				s.logger.Error("failed to rollback transaction", zap.Error(rollbackErr))
			}
		}
	}()

	txRepo := repoproducts.NewRepository(tx, s.logger)
	product, err = txRepo.Create(ctx, p)
	if err != nil {
		s.logger.Error("failed to create product", zap.Error(err),
			zap.String("name", p.Name),
			zap.String("vendor", p.Vendor),
			zap.Float64("price", p.Price.InexactFloat64()))
		return products.Product{}, err
	}

	if ctx.Err() != nil {
		return products.Product{}, ctx.Err()
	}

	if err = s.sqsPublisherRepository.CreateToNotificationsService(ctx, product.ID); err != nil {
		s.logger.Error("failed to send create product msg to notifications services", zap.Error(err),
			zap.String("id", product.ID.String()))
		return products.Product{}, err
	}

	if ctx.Err() != nil {
		return products.Product{}, ctx.Err()
	}

	if err = tx.Commit(); err != nil {
		s.logger.Error("failed to commit transaction", zap.Error(err))
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.logger.Error("failed to rollback after commit error", zap.Error(rollbackErr))
		}
		return products.Product{}, err
	}

	if s.metrics != nil && s.metrics.ProductCreatedCounter != nil {
		s.metrics.ProductCreatedCounter.Inc()
	}

	return product, nil
}

// GetAll returns products with pagination using limit and offset.
func (s Service) GetAll(ctx context.Context, limit, offset uint64) (products.ProductList, error) {
	product, err := s.productsRepository.GetAll(ctx, limit, offset)
	if err != nil {
		s.logger.Error("failed to get all products", zap.Error(err),
			zap.Uint64("limit", limit),
			zap.Uint64("offset", offset))
		return products.ProductList{}, err
	}

	return product, nil
}

// Delete removes a product.
func (s Service) Delete(ctx context.Context, id uuid.UUID) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	if err := s.productsRepository.Delete(ctx, id); err != nil {
		s.logger.Error("failed to delete product", zap.Error(err), zap.String("id", id.String()))
		return err
	}

	if err := s.sqsPublisherRepository.DeleteToNotificationsService(ctx, id); err != nil {
		s.logger.Error("failed to send delete product msg to notifications services", zap.Error(err),
			zap.String("id", id.String()))
		s.logger.Warn("product deleted but notification not sent", zap.String("id", id.String()))
	}

	if s.metrics != nil && s.metrics.ProductDeletedCounter != nil {
		s.metrics.ProductDeletedCounter.Inc()
	}

	return nil
}
