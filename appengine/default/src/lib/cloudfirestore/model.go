package cloudfirestore

import (
	"strings"

	"cloud.google.com/go/firestore"
)

// DocRef ... 個別のドキュメント参照
type DocRef struct {
	CollectionName string `json:"collection_name"`
	DocID          string `json:"doc_id"`
}

// SummaryDocRef ... 最小のドキュメント参照(フロントに返せる)
type SummaryDocRef struct {
	ID   string `json:"id"`
	Path string `json:"path"`
}

// GenerateSummaryDocRef ... 最小のドキュメント参照を取得する
func GenerateSummaryDocRef(docRef *firestore.DocumentRef) *SummaryDocRef {
	if docRef == nil {
		return nil
	}
	dst := &SummaryDocRef{}

	// ID: そのまま
	dst.ID = docRef.ID

	// Path: クライアントSDKの形式に修正
	key := "/documents/"
	if i := strings.Index(docRef.Path, key); i > 0 {
		dst.Path = string(docRef.Path[i+len(key):])
	}
	return dst
}
