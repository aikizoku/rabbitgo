package model

type Sample struct {
	ID        int64  `json:"id" datastore:"-" goon:"id"`
	Name      string `json:"name" datastore:"name,noindex"`
	CreatedAt int64  `json:"created_at" datastore:"created_at,noindex"`
	UpdatedAt int64  `json:"updated_at" datastore:"updated_at,noindex"`
}

type ArticleList struct {
	ID int64 `datastore:"-" goon:"id"`

	Type        string `datastore:"type"`
	Title       string `datastore:"title"`
	Description string `datastore:"description,noindex"`
	PublishedAt int64  `datastore:"published_at"`
}
