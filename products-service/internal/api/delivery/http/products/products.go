package products

import (
	"github.com/at-kh/guru-apps-test-services/products-service/internal/api/delivery"
	dproducts "github.com/at-kh/guru-apps-test-services/products-service/internal/api/domain/products"
	"github.com/at-kh/guru-apps-test-services/products-service/internal/api/services"
	"github.com/at-kh/guru-apps-test-services/products-service/pkg/errs"
	"github.com/at-kh/guru-apps-test-services/products-service/pkg/responder"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"go.uber.org/zap"
)

var _ delivery.ProductsHTTPHandler = &Handler{}

type (
	// Handler defines a Handler for HTTP requests for checking processing voice.
	Handler struct {
		responder.Responder

		service services.ProductsService
		log     *zap.Logger
	}
)

// NewHandler - create new handler.
func NewHandler(responder responder.Responder, service services.ProductsService, log *zap.Logger) *Handler {
	return &Handler{
		Responder: responder,
		service:   service,
		log:       log.With(zap.String("http_handler", "products")),
	}
}

// Create - create new product.
func (h Handler) Create(ctx *fiber.Ctx) error {
	var req createProductRequest
	if err := ctx.BodyParser(&req); err != nil {
		return errs.BadRequest{Cause: "invalid JSON body"}
	}

	if errsList := req.Validate(); len(errsList) != 0 {
		return errs.FieldsValidation{Errors: errsList}
	}

	product, err := h.service.Create(ctx.Context(), req.toDomain())
	if err != nil {
		return err
	}

	return h.Respond(ctx, fiber.StatusCreated, fromDomain(product))
}

// GetAll - get all products:
//   - GET /products
//   - GET /products?limit=2&offset=1
func (h Handler) GetAll(ctx *fiber.Ctx) error {
	limit := uint64(ctx.QueryInt("limit", dproducts.DefaultLimit))
	if limit == 0 {
		limit = dproducts.DefaultLimit
	}
	if limit > dproducts.MaxLimit {
		limit = dproducts.MaxLimit
	}

	offsetInt := ctx.QueryInt("offset", dproducts.DefaultOffset)
	if offsetInt < 0 {
		offsetInt = dproducts.DefaultOffset
	}
	offset := uint64(offsetInt)

	list, err := h.service.GetAll(ctx.Context(), limit, offset)
	if err != nil {
		return err
	}

	return h.Respond(ctx, fiber.StatusOK, fromDomainList(list, limit, offset))
}

// Delete - delete product by id.
func (h Handler) Delete(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	if idStr == "" {
		return errs.BadRequest{Cause: "id is required"}
	}

	productID, err := uuid.Parse(idStr)
	if err != nil {
		return err
	}

	if err = h.service.Delete(ctx.Context(), productID); err != nil {
		return err
	}

	return h.RespondEmpty(ctx, fiber.StatusNoContent)
}
