package model

import (
	"time"
)

// Sample ... サンプルモデル
type Sample struct {
	ID        int64     `json:"id" datastore:"-" boom:"id"`
	Category  string    `json:"category"`
	Name      string    `json:"name" datastore:",noindex"`
	Enabled   bool      `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at" datastore:"-"`
}
