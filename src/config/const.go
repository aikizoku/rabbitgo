package config

import "time"

// ContextKey ... ContextKeyの型定義
type ContextKey string

const (
	// HTTPRequestTimeout ... デフォルトのHTTPリクエストタイムアウト
	HTTPRequestTimeout time.Duration = 20

	// FirebaseCredentialFilePath ... FirebaseSDKのCredentialFileのPath
	FirebaseCredentialFilePath string = "./firebase_credentials.json"
)
