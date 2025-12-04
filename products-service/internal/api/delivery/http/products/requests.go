package products

import (
	"github.com/at-kh/guru-apps-test-services/products-service/internal/api/domain/products"
	"github.com/shopspring/decimal"
)

//go:generate go-validator

type (
	// createProductRequest - request model for creation.
	createProductRequest struct {
		Name        string          `json:"name" valid:"required,max=255"`
		Vendor      string          `json:"vendor" valid:"required,max=255"`
		Description string          `json:"description" valid:"max=10000"`
		Price       decimal.Decimal `json:"price" valid:"min=0,max=1000000"`
	}
)

// toDomain converts request model to domain model.
func (r createProductRequest) toDomain() products.Product {
	return products.Product{
		Name:        r.Name,
		Vendor:      r.Vendor,
		Description: r.Description,
		Price:       r.Price,
	}
}
