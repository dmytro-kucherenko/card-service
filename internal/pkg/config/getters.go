package config

import (
	"fmt"
)

const (
	appPort     = 8000
	appProtocol = "http"
	appHostName = "localhost"
	appBasePath = "/api"
)

func AppPort() uint16 {
	if config.AppPort == 0 {
		return appPort
	}

	return config.AppPort
}

func AppProtocol() string {
	if config.AppProtocol == "" {
		return appProtocol
	}

	return config.AppProtocol
}

func AppHost() string {
	if config.AppHost == "" {
		return fmt.Sprintf("%s:%v", appHostName, AppPort())
	}

	return config.AppHost
}

func AppBasePath() string {
	if config.AppBasePath == "" {
		return appBasePath
	}

	return config.AppBasePath
}
