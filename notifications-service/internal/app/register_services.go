package app

import "github.com/at-kh/guru-apps-test-services/notifications-service/internal/api/services/notifications"

// registerServices register services in-app struct.
func (a *App) registerServices() {
	a.notificationsService = notifications.NewService(a.logger)
}
