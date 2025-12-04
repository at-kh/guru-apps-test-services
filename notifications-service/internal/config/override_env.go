package config

import (
	"os"

	"github.com/joho/godotenv"
)

// overrideEnv - override config fields from env.
func (cfg *Config) overrideEnv() error {
	_ = godotenv.Load()

	override := func(env string, dst *string) {
		if v := os.Getenv(env); v != "" {
			*dst = v
		}
	}

	override("HTTP_ADDRESS", &cfg.Delivery.HTTPServer.ListenAddress)
	override("SQS_URL", &cfg.Delivery.Broker.URL)
	override("SQS_REGION", &cfg.Delivery.Broker.Region)

	return nil
}
