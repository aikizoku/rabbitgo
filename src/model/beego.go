package model

// Beego ... サンプルモデル
type Beego struct {
	ID        int64       `json:"id" datastore:"-" goon:"id"`
	Name      string      `json:"name" datastore:",noindex"`
	Types     []BeegoType `json:"types" datastore:"-"`
	Enabled   bool        `json:"-"`
	CreatedAt int64       `json:"created_at" datastore:",noindex"`
	UpdatedAt int64       `json:"updated_at" datastore:",noindex"`
}

// BeegoType ... サンプルモデルの種類
type BeegoType struct {
	ID        int64  `json:"id" datastore:"-" goon:"id"`
	Name      string `json:"name" datastore:",noindex"`
	CreatedAt int64  `json:"created_at" datastore:",noindex"`
	UpdatedAt int64  `json:"updated_at" datastore:",noindex"`
}
