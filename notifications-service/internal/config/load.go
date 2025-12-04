package config

import (
	"errors"
	"fmt"
	"os"
	"reflect"

	"gopkg.in/yaml.v3"
)

// InitConfig - init configuration from yaml and env.
func InitConfig(configPath string) (*Config, error) {
	cfg := &Config{}

	if err := loadYAML(configPath, cfg); err != nil {
		return nil, fmt.Errorf("failed to load yaml: %w", err)
	}

	if err := cfg.overrideEnv(); err != nil {
		return nil, fmt.Errorf("failed to load env: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("configuration validation error: %v", err)
	}

	return cfg, nil
}

// loadYAML - load config from yaml.
func loadYAML(path string, dest any) error {
	if reflect.TypeOf(dest).Kind() != reflect.Ptr {
		return errors.New("config destination must be pointer")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("configuration file not found: %s", path)
		}

		return fmt.Errorf("failed to read configuration file %s: %w", path, err)
	}

	if len(data) == 0 {
		return fmt.Errorf("configuration file is empty: %s", path)
	}

	if err = yaml.Unmarshal(data, dest); err != nil {
		return fmt.Errorf("failed to parse YAML configuration file %s: %w", path, err)
	}

	return nil
}
