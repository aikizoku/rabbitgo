package main

import (
	"context"

	"github.com/aikizoku/rabbitgo/command/lib"
	"github.com/aikizoku/rabbitgo/command/seed/content"
)

func main() {
	env := lib.GetEnv()
	d := lib.NewDependency(env)
	c := NewContent(d)
	ctx := context.Background()

	c.Sample.Generate(ctx)
}

type Content struct {
	Sample *content.Sample
}

func NewContent(d *lib.Dependency) *Content {
	return &Content{
		content.NewSample(d),
	}
}
