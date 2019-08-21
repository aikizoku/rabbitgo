package p

// Param ... PubSubから受け取るパラメータ
type Param struct {
	SourceID  string    `json:"source_id"`
	SourceURL string    `json:"source_url"`
	OutPath   string    `json:"out_path"`
	DocRefs   []*DocRef `json:"doc_refs"`
	FieldName string    `json:"field_name"`
}

// Preset ... 画像変換のプリセット
type Preset struct {
	Name   string `json:"name"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// DocRef ... 個別のドキュメント参照
type DocRef struct {
	CollectionName string `json:"collection_name"`
	DocID          string `json:"doc_id"`
}

// Object ... 画像オブジェクト
type Object struct {
	ID            string           `firestore:"id"`
	URL           string           `firestore:"url"`
	DominantColor string           `firestore:"dominant_color"`
	Sizes         map[string]*Size `firestore:"sizes"`
}

// Size ... サイズ毎の画像情報
type Size struct {
	URL    string `firestore:"url"`
	Width  int    `firestore:"width"`
	Height int    `firestore:"height"`
}
