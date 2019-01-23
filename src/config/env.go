package config

import (
	"os"
)

// IsEnvLocal ... 現在の環境がローカルか判定する
func IsEnvLocal() bool {
	return os.Getenv("ENV") == "local"
}

// IsEnvStaging ... 現在の環境がステージングか判定する
func IsEnvStaging() bool {
	return os.Getenv("ENV") == "staging"
}

// IsEnvProduction ... 現在の環境が本番か判定する
func IsEnvProduction() bool {
	return os.Getenv("ENV") == "production"
}
