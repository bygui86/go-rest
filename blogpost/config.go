package blogpost

import (
	"github.com/bygui86/go-rest/logging"
	"github.com/bygui86/go-rest/utils"
)

const (
	restHostEnvVar = "BLOGPOSTS_REST_HOST"
	restPortEnvVar = "BLOGPOSTS_REST_PORT"

	restHostEnvVarDefault = "localhost"
	restPortEnvVarDefault = 8080
)

func loadConfig() *config {
	logging.Log.Debug("Load configurations")
	return &config{
		RestHost: utils.GetStringEnv(restHostEnvVar, restHostEnvVarDefault),
		RestPort: utils.GetIntEnv(restPortEnvVar, restPortEnvVarDefault),
	}
}
