package main

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"

	"github.com/aikizoku/rabbitgo/appengine/src/lib/deploy"
)

// Environment ... 環境変数
type Environment struct {
	Port            int    `envconfig:"PORT"                           default:"8080"`
	Deploy          string `envconfig:"DEPLOY"                         required:"true"`
	ProjectID       string `envconfig:"PROJECT_ID"                     required:"true"`
	LocationID      string `envconfig:"LOCATION_ID"                    default:"asia-northeast1"`
	CredentialsPath string `envconfig:"GOOGLE_APPLICATION_CREDENTIALS" required:"true"`
	MinLogSeverity  string `envconfig:"MIN_LOG_SEVERITY"               required:"true"`
}

// Get ... 環境変数を取得する
func (e *Environment) Get() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	err = envconfig.Process("", e)
	if err != nil {
		panic(err)
	}
	if deploy.IsLocal() {

	}
}
