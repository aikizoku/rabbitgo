package images

import "context"

// Repository ... リポジトリ
type Repository interface {
	UpdateObjects(ctx context.Context, key string, objects []*Object) error
}
