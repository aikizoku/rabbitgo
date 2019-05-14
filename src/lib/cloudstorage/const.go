package cloudstorage

const (
	// BaseURL ... GCSのURL
	BaseURL string = "https://storage.googleapis.com"
	// ChunkSize ... アップロード時の分割サイズ（メモリ不足になったら調整する）
	ChunkSize int = 200
)
