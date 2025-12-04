package sqs_publisher

import "github.com/google/uuid"

// Event types for notifications services.
const (
	eventTypeCreateProduct = "create_product"
	eventTypeDeleteProduct = "delete_product"
)

// message - message for SQS.
type message struct {
	EventType string    `json:"event_type"`
	ProductID uuid.UUID `json:"product_id"`
}
