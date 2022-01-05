package app

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/rabee-inc/go-pkg/environment"

	"github.com/aikizoku/rabbitgo/appengine/api/src/config"
)

// Environment ... 環境変数
type Environment struct {
	Port              int    `envconfig:"PORT"                default:"8080"`
	Deploy            string `envconfig:"DEPLOY"              default:"local"`
	ProjectID         string `envconfig:"PROJECT_ID"          required:"true"`
	LocationID        string `envconfig:"LOCATION_ID"         default:"asia-northeast1"`
	MinLogSeverity    string `envconfig:"MIN_LOG_SEVERITY"    required:"true"`
	InternalAuthToken string `envconfig:"INTERNAL_AUTH_TOKEN" required:"true"`
}

// Get ... 環境変数を取得する
func (e *Environment) Get() {
	environment.Load(config.GetFilePath("env.yaml"))
	err := envconfig.Process("", e)
	if err != nil {
		panic(err)
	}
}
