package services

import (
	"github.com/google/uuid"
)

// Notifications - interface for notifications services.
type Notifications interface {
	Create(id uuid.UUID)
	Delete(id uuid.UUID)
}
