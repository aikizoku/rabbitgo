package deployed

import (
	"os"
)

// IsLocal ... 現在の環境がローカルか判定する
func IsLocal() bool {
	return os.Getenv("DEPLOY") == "local"
}

// IsStaging ... 現在の環境がステージングか判定する
func IsStaging() bool {
	return os.Getenv("DEPLOY") == "staging"
}

// IsProduction ... 現在の環境が本番か判定する
func IsProduction() bool {
	return os.Getenv("DEPLOY") == "production"
}
