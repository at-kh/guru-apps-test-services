package products

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type (
	// Product struct represents a product entity.
	Product struct {
		ID          uuid.UUID
		Name        string
		Vendor      string
		Description string
		Price       decimal.Decimal
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}

	// Products describe a list of Products.
	Products []Product

	// ProductList describes a list of Products with their total number.
	ProductList struct {
		Total    uint64
		Products Products
	}
)
