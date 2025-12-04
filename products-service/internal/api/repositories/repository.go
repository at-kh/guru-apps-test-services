package repositories

import (
	"context"

	"github.com/at-kh/guru-apps-test-services/products-service/internal/api/domain/products"
	"github.com/google/uuid"
)

type (
	// ProductsRepository defines the interface for product repositories.
	ProductsRepository interface {
		Create(ctx context.Context, e products.Product) (products.Product, error)
		GetAll(ctx context.Context, limit, offset uint64) (products.ProductList, error)
		Delete(ctx context.Context, id uuid.UUID) error
	}

	// SQSPublisherRepository defines the interface for SQS publisher.
	SQSPublisherRepository interface {
		CreateToNotificationsService(ctx context.Context, id uuid.UUID) error
		DeleteToNotificationsService(ctx context.Context, id uuid.UUID) error
	}
)
