package cloudstorage

import "time"

// CacheMode ... キャッシュ設定
type CacheMode struct {
	Disabled bool
	Expire   time.Duration
}

// UploadResponse ... アップロードのレスポンス
type UploadResponse struct {
	URL string
}
