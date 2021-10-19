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

// Content ... 実処理
type Content struct {
	Sample *content.Sample
}

// NewContent ... 実処理を取得する
func NewContent(d *lib.Dependency) *Content {
	return &Content{
		Sample: content.NewSample(d),
	}
}
