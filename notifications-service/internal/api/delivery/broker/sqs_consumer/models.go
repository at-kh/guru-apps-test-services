package sqs_consumer

import "github.com/google/uuid"

// message - defines message struct for SQS consumer, supported event types:
//   - create_product
//   - delete_product
type message struct {
	EventType string    `json:"event_type"`
	ProductID uuid.UUID `json:"product_id"`
}
