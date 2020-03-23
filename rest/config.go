package rest

import (
	"github.com/bygui86/go-rest/logging"
	"github.com/bygui86/go-rest/utils"
)

const (
	restHostEnvVar = "REST_HOST"
	restPortEnvVar = "REST_PORT"

	restHostEnvVarDefault = "localhost"
	restPortEnvVarDefault = 8080
)

type config struct {
	RestHost string
	RestPort int
}

func loadConfig() *config {
	logging.Log.Debug("Load REST configurations")
	return &config{
		RestHost: utils.GetStringEnv(restHostEnvVar, restHostEnvVarDefault),
		RestPort: utils.GetIntEnv(restPortEnvVar, restPortEnvVarDefault),
	}
}
