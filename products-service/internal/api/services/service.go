package services

import (
	"context"

	"github.com/at-kh/guru-apps-test-services/products-service/internal/api/domain/products"
	"github.com/google/uuid"
)

type (
	// ProductsService defines the interface for product services.
	ProductsService interface {
		Create(ctx context.Context, e products.Product) (products.Product, error)
		GetAll(ctx context.Context, page, limit uint64) (products.ProductList, error)
		Delete(ctx context.Context, id uuid.UUID) error
	}
)
