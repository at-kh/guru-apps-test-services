package products

import (
	"database/sql"
	"time"

	"github.com/at-kh/guru-apps-test-services/products-service/internal/api/domain/products"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type (
	// dbProduct - defines a product in the database.
	dbProduct struct {
		ID          uuid.UUID       `db:"id"`
		Name        string          `db:"name"`
		Vendor      string          `db:"vendor"`
		Description sql.NullString  `db:"description"`
		Price       decimal.Decimal `db:"price"`
		CreatedAt   time.Time       `db:"created_at"`
		UpdatedAt   time.Time       `db:"updated_at"`
	}
)

// toDomain converts dbProduct -> Product
func (d dbProduct) toDomain() products.Product {
	desc := ""
	if d.Description.Valid {
		desc = d.Description.String
	}

	return products.Product{
		ID:          d.ID,
		Name:        d.Name,
		Vendor:      d.Vendor,
		Description: desc,
		Price:       d.Price,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}
