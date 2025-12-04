package app

import (
	"github.com/at-kh/guru-apps-test-services/notifications-service/internal/api/delivery/broker/sqs_consumer"
)

// registerBrokerHandlers initializes the broker handlers.
func (a *App) registerBrokerHandlers() {
	a.sqsConsumerHandler = sqs_consumer.NewHandler(a.notificationsService, a.logger)
}
