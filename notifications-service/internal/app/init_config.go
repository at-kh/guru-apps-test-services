package app

import (
	"github.com/at-kh/guru-apps-test-services/notifications-service/internal/config"
	"github.com/gofiber/fiber/v2/log"
)

// initConfig - init config from yaml and env with validation.
func (a *App) initConfig() {
	cfg, err := config.InitConfig(config.DefaultPath)
	if err != nil {
		log.Fatal("config error: %v", err)
	}

	a.cfg = cfg
}
