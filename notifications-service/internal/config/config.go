package config

import (
	"time"
)

//go:generate go-validator

// DefaultPath - default path for config.
const DefaultPath = "./cmd/config.yaml"

// Config defines the properties of the application configuration.
type (
	Config struct {
		Delivery Delivery `yaml:"delivery" valid:"check,deep"`
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

		RetryDelay          time.Duration `yaml:"retry-delay" valid:"required"`
		DeleteTimeout       time.Duration `yaml:"delete-timeout" valid:"required"`
		HandlerTimeout      time.Duration `yaml:"handler-timeout" valid:"required"`
		MaxNumberOfMessages int32         `yaml:"max-number-of-messages" valid:"required"`
		WaitTimeSeconds     int32         `yaml:"wait-time-seconds" valid:"required"`
	}
)
