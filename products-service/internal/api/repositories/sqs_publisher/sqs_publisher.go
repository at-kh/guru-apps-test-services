package sqs_publisher

import (
	"context"

	"github.com/at-kh/guru-apps-test-services/products-service/internal/api/repositories"
	"github.com/at-kh/guru-apps-test-services/products-service/pkg/errs"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

var _ repositories.SQSPublisherRepository = &Repository{}

// Repository implements repositories interface for publishing messages to SQS.
type Repository struct {
	client       *sqs.Client
	baseQueueURL string
	logger       *zap.Logger
}

// NewRepository creates a new repositories.
func NewRepository(client *sqs.Client, baseQueueURL string, logger *zap.Logger) repositories.SQSPublisherRepository {
	return &Repository{
		client:       client,
		baseQueueURL: baseQueueURL,
		logger:       logger.With(zap.String("repositories", "sqs_publisher")),
	}
}

// sendEvent sends an event to SQS.
func (r Repository) sendEvent(ctx context.Context, msg message) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return errs.Internal{Cause: "failed to marshal SQS event: " + err.Error()}
	}

	input := &sqs.SendMessageInput{
		QueueUrl:    aws.String(r.baseQueueURL),
		MessageBody: aws.String(string(data)),
	}

	if _, err = r.client.SendMessage(ctx, input); err != nil {
		return errs.Internal{Cause: "failed to send message to SQS: " + err.Error()}
	}

	r.logger.Debug("event sent to SQS",
		zap.String("event_type", msg.EventType),
		zap.String("product_id", msg.ProductID.String()),
	)

	return nil
}

// CreateToNotificationsService - send event "created product".
func (r Repository) CreateToNotificationsService(ctx context.Context, id uuid.UUID) error {
	return r.sendEvent(ctx, message{
		EventType: eventTypeCreateProduct,
		ProductID: id,
	})
}

// DeleteToNotificationsService - send event "deleted product".
func (r Repository) DeleteToNotificationsService(ctx context.Context, id uuid.UUID) error {
	return r.sendEvent(ctx, message{
		EventType: eventTypeDeleteProduct,
		ProductID: id,
	})
}
