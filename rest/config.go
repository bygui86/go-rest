package rest

import (
	"github.com/bygui86/go-save-memory/http-server/logging"
	"github.com/bygui86/go-save-memory/http-server/utils"
)

const (
	restHostEnvVar = "REST_HOST"
	restPortEnvVar = "REST_PORT"

	restHostDefault = "localhost"
	restPortDefault = 8080
)

type Config struct {
	RestHost string
	RestPort int
}

func loadConfig() *Config {
	logging.Log.Debug("Load REST configurations")
	return &Config{
		RestHost: utils.GetStringEnv(restHostEnvVar, restHostDefault),
		RestPort: utils.GetIntEnv(restPortEnvVar, restPortDefault),
	}
}
