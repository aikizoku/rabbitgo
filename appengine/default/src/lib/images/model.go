package images

// Object ... 画像オブジェクト
type Object struct {
	ID            string           `firestore:"id"             json:"id"`
	URL           string           `firestore:"url"            json:"url"`
	DominantColor string           `firestore:"dominant_color" json:"dominant_color"`
	Sizes         map[string]*Size `firestore:"sizes"          json:"sizes"`
	IsDefault     bool             `firestore:"is_default"     json:"is_default"`
	Ref           string           `firestore:"ref"            json:"-"`
}

// Size ... サイズ
type Size struct {
	URL    string `firestore:"url"    json:"url"`
	Width  int    `firestore:"width"  json:"width"`
	Height int    `firestore:"height" json:"height"`
}

// ConvRequest ... 画像変換リクエスト
type ConvRequest struct {
	Key         string   `json:"key"`
	SourceID    string   `json:"source_id"`
	SourceURLs  []string `json:"source_urls"`
	DstFilePath string   `json:"dst_file_path"`
}
