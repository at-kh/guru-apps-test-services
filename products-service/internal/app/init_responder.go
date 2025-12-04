package app

import "github.com/at-kh/guru-apps-test-services/products-service/pkg/responder"

// initResponder initializes responder.
func (a *App) initResponder() {
	a.responder = responder.New()
}
