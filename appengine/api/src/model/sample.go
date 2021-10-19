package model

import (
	"cloud.google.com/go/firestore"
	"github.com/rabee-inc/go-pkg/images"
)

// Sample ... サンプルモデル
type Sample struct {
	ID           string                 `firestore:"-" cloudfirestore:"id"`
	Ref          *firestore.DocumentRef `firestore:"-" cloudfirestore:"ref"`
	Category     string                 `firestore:"category"`
	Name         string                 `firestore:"name"`
	ProfileImage *images.Object         `firestore:"profile_image"`
	Disabled     bool                   `firestore:"disabled"`
	CreatedAt    int64                  `firestore:"created_at"`
	UpdatedAt    int64                  `firestore:"updated_at"`
}
