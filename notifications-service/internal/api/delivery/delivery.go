package delivery

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

// HTTP handlers
type (
	// HealthHTTPHandler - describes an interface for work with services health over HTTP.
	HealthHTTPHandler interface {
		// Health - handler for getting meta information endpoint.
		Health(ctx *fiber.Ctx) error
	}
)

// Brokers handlers
type (
	// SQSConsumerHandler - describes an interface for work with SQS consumer over HTTP.
	SQSConsumerHandler interface {
		CreateNotification(ctx context.Context, body []byte) error
		DeleteNotification(ctx context.Context, body []byte) error
	}
)
