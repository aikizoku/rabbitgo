package model

type Sample struct {
	ID        int64  `json:"id" datastore:"_" goon:"id"`
	Name      string `json:"name" datastore:"name,noindex"`
	CreatedAt int64  `json:"created_at" datastore:"created_at,noindex"`
	UpdatedAt int64  `json:"updated_at" datastore:"updated_at,noindex"`
}
