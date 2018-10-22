package model

// Sample ... サンプルモデル
type Sample struct {
	ID        int64  `json:"id" datastore:"-" goon:"id"`
	Category  string `json:"category"`
	Name      string `json:"name" datastore:",noindex"`
	Enabled   bool   `json:"-"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at" datastore:"-"`
}
