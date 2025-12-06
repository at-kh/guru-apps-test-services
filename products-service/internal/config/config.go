package config

import (
	"time"
)

//go:generate go-validator

// DefaultPath - default path for config.
const DefaultPath = "./cmd/config.yaml"

type (
	// Config defines the properties of the application configuration.
	Config struct {
		Delivery Delivery `yaml:"delivery" valid:"check,deep"`
		Storage  Storage  `yaml:"storage"  valid:"check,deep"`
	}

	// Delivery defines API server configuration.
	Delivery struct {
		HTTPServer HTTPServer `yaml:"http-server" valid:"check,deep"`
		Broker     Broker     `yaml:"broker" valid:"check,deep"`
	}

	// HTTPServer defines the HTTP section of the API server configuration.
	HTTPServer struct {
		ListenAddress      string        `yaml:"listen-address"  valid:"required"`
		ReadTimeout        time.Duration `yaml:"read-timeout"    valid:"required"`
		WriteTimeout       time.Duration `yaml:"write-timeout"   valid:"required"`
		GracefulTimeout    time.Duration `yaml:"graceful-timeout" valid:"required"`
		BodySizeLimitBytes int           `yaml:"body-size-limit" valid:"required"`
	}

	// Broker defines the message queue section of the API server configuration.
	Broker struct {
		URL    string `valid:"check,deep"`
		Region string `valid:"check,deep"`
	}

	// Storage defines the storage section of the API server configuration.
	Storage struct {
		Postgres Postgres `yaml:"postgres" valid:"check,deep"`
	}

	// Postgres defines the Postgres section of the API server configuration.
	Postgres struct {
		DSN                string        `yaml:"dsn"                   valid:"required"`
		Driver             string        `yaml:"driver"                valid:"required"`
		Dialect            string        `yaml:"dialect"               valid:"required"`
		MigrationDirectory string        `yaml:"migration-directory"   valid:"required"`
		MigrationDirection string        `yaml:"migration-direction"   valid:"required"`
		ConnMaxLifetime    time.Duration `yaml:"conn-max-lifetime"`
		RetryDelay         time.Duration `yaml:"retry-delay"           valid:"required"`
		QueryTimeout       time.Duration `yaml:"query-timeout"`
		ConnMaxIdleNum     int           `yaml:"conn-max-idle-num"     valid:"required"`
		ConnMaxOpenNum     int           `yaml:"conn-max-open-num"     valid:"required"`
		MaxRetries         int           `yaml:"max-retries"           valid:"required,min=0"`
		AutoMigrate        bool          `yaml:"auto-migrate"`
	}
)
