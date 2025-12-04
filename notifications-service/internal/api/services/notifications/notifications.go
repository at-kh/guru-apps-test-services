package notifications

import (
	"github.com/at-kh/guru-apps-test-services/notifications-service/internal/api/services"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

var _ services.Notifications = &Service{}

// Service - defines services struct.
type Service struct {
	logger *zap.Logger
}

// NewService constructor.
func NewService(logger *zap.Logger) *Service {
	return &Service{
		logger: logger.With(zap.String("service", "notifications")),
	}
}

// Create - log info about created product.
func (s Service) Create(id uuid.UUID) {
	s.logger.Info("➕ product created", zap.String("id", id.String()))
}

// Delete - log info about deleted product.
func (s Service) Delete(id uuid.UUID) {
	s.logger.Info("➖ product deleted", zap.String("id", id.String()))
}
