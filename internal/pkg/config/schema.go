package config

import (
	"github.com/dmytro-kucherenko/card-service/internal/pkg/log"
)

type schema struct {
	AppPort     uint16 `mapstructure:"APP_PORT"`
	AppProtocol string `mapstructure:"APP_PROTOCOL" validate:"omitempty,oneof=http https"`
	AppHost     string `mapstructure:"APP_HOST"`
	AppBasePath string `mapstructure:"APP_BASE_PATH" validate:"omitempty,startswith=/"`
}

var config *schema

func init() {
	logger := log.NewConsole("Config")
	config = new(schema)

	if err := load(".env", config); err != nil {
		logger.Fatal(err)
	}

	logger.Info("Environmental variables loaded")
}
