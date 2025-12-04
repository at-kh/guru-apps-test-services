package sqs_consumer

import (
	"context"

	"github.com/at-kh/guru-apps-test-services/notifications-service/internal/api/delivery"
	"github.com/at-kh/guru-apps-test-services/notifications-service/internal/api/services"
	"github.com/at-kh/guru-apps-test-services/notifications-service/pkg/errs"
	"github.com/goccy/go-json"
	"go.uber.org/zap"
)

var _ delivery.SQSConsumerHandler = &Handler{}

type Handler struct {
	service services.Notifications
	logger  *zap.Logger
}

// NewHandler - creates a new SQS consumer handler.
func NewHandler(service services.Notifications, logger *zap.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger.With(zap.String("broker", "sqs_consumer")),
	}
}

// CreateNotification - log "create product" by SQS message.
func (h Handler) CreateNotification(_ context.Context, body []byte) error {
	var msg message
	if err := json.Unmarshal(body, &msg); err != nil {
		return errs.Internal{Cause: err.Error()}
	}

	h.logger.Debug("processing created product", zap.String("product_id", msg.ProductID.String()))

	h.service.Create(msg.ProductID)

	return nil
}

// DeleteNotification - log "delete product" by SQS message.
func (h Handler) DeleteNotification(_ context.Context, body []byte) error {
	var msg message
	if err := json.Unmarshal(body, &msg); err != nil {
		return errs.Internal{Cause: err.Error()}
	}

	h.logger.Debug("processing deleted product", zap.String("product_id", msg.ProductID.String()))

	h.service.Delete(msg.ProductID)

	return nil
}
