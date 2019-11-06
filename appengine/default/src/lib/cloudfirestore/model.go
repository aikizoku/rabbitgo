package cloudfirestore

// DocRef ... 個別のドキュメント参照
type DocRef struct {
	CollectionName string `json:"collection_name"`
	DocID          string `json:"doc_id"`
}
