package products

import (
	"context"
	"strings"

	"github.com/at-kh/guru-apps-test-services/products-service/internal/api/domain/products"
	"github.com/at-kh/guru-apps-test-services/products-service/internal/api/repositories"
	"github.com/at-kh/guru-apps-test-services/products-service/pkg/errs"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var _ repositories.ProductsRepository = &Repository{}

type (
	// Repository - defines a repositories.
	Repository struct {
		db     sqlx.ExtContext
		logger *zap.Logger
	}
)

// NewRepository creates a new repositories.
func NewRepository(db sqlx.ExtContext, logger *zap.Logger) repositories.ProductsRepository {
	return &Repository{db: db, logger: logger.With(zap.String("repositories", "products"))}
}

// Create creates a new product.
func (r *Repository) Create(ctx context.Context, e products.Product) (products.Product, error) {
	if ctx.Err() != nil {
		return products.Product{}, ctx.Err()
	}

	query := `
	INSERT INTO products (name, vendor, description, price)
	VALUES ($1, $2, $3, $4)
	RETURNING id, name, vendor, description, price, created_at, updated_at;
	`

	var dbp dbProduct
	if err := sqlx.GetContext(ctx, r.db, &dbp, query, e.Name, e.Vendor, e.Description, e.Price); err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return products.Product{}, errs.Conflict{What: "product already exists"}
		}

		return products.Product{}, errs.Internal{Cause: err.Error()}
	}

	return dbp.toDomain(), nil
}

// GetAll returns products with pagination using limit and offset.
func (r *Repository) GetAll(ctx context.Context, limit, offset uint64) (products.ProductList, error) {
	if ctx.Err() != nil {
		return products.ProductList{}, ctx.Err()
	}

	if limit == 0 {
		limit = products.DefaultLimit
	}

	query := `
	SELECT id, name, vendor, description, price, created_at, updated_at
	FROM products
	ORDER BY created_at DESC
	LIMIT $1 OFFSET $2;
	`

	var dbItems []dbProduct
	if err := sqlx.SelectContext(ctx, r.db, &dbItems, query, limit, offset); err != nil {
		return products.ProductList{}, errs.Internal{Cause: err.Error()}
	}

	items := make([]products.Product, len(dbItems))
	for i, dbp := range dbItems {
		items[i] = dbp.toDomain()
	}

	// get total count of entities
	var total uint64
	if err := sqlx.GetContext(ctx, r.db, &total, "SELECT COUNT(*) FROM products"); err != nil {
		return products.ProductList{}, errs.Internal{Cause: err.Error()}
	}

	return products.ProductList{
		Total:    total,
		Products: items,
	}, nil
}

// Delete removes a product.
func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	query := `DELETE FROM products WHERE id = $1`

	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return errs.Internal{Cause: err.Error()}
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return errs.Internal{Cause: err.Error()}
	}
	if affected == 0 {
		return errs.NotFound{What: "product"}
	}

	return nil
}
