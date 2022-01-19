package content

import (
	"context"

	"github.com/aikizoku/rabbitgo/command/lib"
)

type Sample struct {
	d *lib.Dependency
}

func (m *Sample) Generate(ctx context.Context) {
}

func NewSample(d *lib.Dependency) *Sample {
	return &Sample{
		d: d,
	}
}
