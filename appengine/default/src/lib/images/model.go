package images

import "github.com/aikizoku/rabbitgo/appengine/default/src/lib/cloudfirestore"

// Object ... 画像オブジェクト
type Object struct {
	ID            string           `firestore:"id"             json:"id"`
	URL           string           `firestore:"url"            json:"url"`
	DominantColor string           `firestore:"dominant_color" json:"dominant_color"`
	Sizes         map[string]*Size `firestore:"sizes"          json:"sizes"`
}

// Size ... サイズ
type Size struct {
	URL    string `firestore:"url"    json:"url"`
	Width  int    `firestore:"width"  json:"width"`
	Height int    `firestore:"height" json:"height"`
}

// ConvRequest ... 画像変換リクエスト
type ConvRequest struct {
	SourceID  string                   `json:"source_id"`
	SourceURL string                   `json:"source_url"`
	OutPath   string                   `json:"out_path"`
	DocRefs   []*cloudfirestore.DocRef `json:"doc_refs"`
	FieldName string                   `json:"field_name"`
}
