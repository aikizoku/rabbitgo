package content

import (
	"context"

	"github.com/aikizoku/rabbitgo/command/lib"
)

// Sample ... サンプルのシードデータ
type Sample struct {
	d *lib.Dependency
}

// Generate ... シードデータを作成する
func (m *Sample) Generate(ctx context.Context) {
}

// NewSample ...
func NewSample(d *lib.Dependency) *Sample {
	return &Sample{
		d: d,
	}
}
