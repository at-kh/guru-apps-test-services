package products

import (
	"time"

	"github.com/at-kh/guru-apps-test-services/products-service/internal/api/domain/products"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type (
	// paginationResponse is a model to store pagination parameters.
	paginationResponse struct {
		Offset uint64 `query:"offset" json:"offset"`
		Limit  uint64 `query:"limit" json:"limit"`
		Total  uint64 `query:"total" json:"total"`
	}

	// productResponse is a response for a product.
	productResponse struct {
		ID          uuid.UUID       `json:"id"`
		Name        string          `json:"name"`
		Vendor      string          `json:"vendor"`
		Description string          `json:"description"`
		Price       decimal.Decimal `json:"price"`
		CreatedAt   time.Time       `json:"created_at"`
		UpdatedAt   time.Time       `json:"updated_at"`
	}

	// productListResponse is a response for a list of products.
	productListResponse struct {
		Pagination paginationResponse `json:"pagination"`
		Products   []productResponse  `json:"products"`
	}
)

// fromDomain converts domain model to response model.
func fromDomain(p products.Product) productResponse {
	return productResponse{
		ID:          p.ID,
		Name:        p.Name,
		Vendor:      p.Vendor,
		Description: p.Description,
		Price:       p.Price,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

// fromDomainList converts domain model to response model.
func fromDomainList(list products.ProductList, limit, offset uint64) productListResponse {
	result := make([]productResponse, 0, list.Total)
	for _, msg := range list.Products {
		result = append(result, fromDomain(msg))
	}

	return productListResponse{
		Pagination: paginationResponse{
			Offset: offset,
			Limit:  limit,
			Total:  list.Total,
		},
		Products: result,
	}
}
