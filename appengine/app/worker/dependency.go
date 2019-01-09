package main

import (
	"github.com/aikizoku/gocci/src/handler/worker"
)

// Dependency ... 依存性
type Dependency struct {
	AdminHandler  *worker.AdminHandler
	SampleHandler *worker.SampleHandler
}

// Inject ... 依存性を注入する
func (d *Dependency) Inject() {
	// Handler
	d.AdminHandler = &worker.AdminHandler{}
	d.SampleHandler = &worker.SampleHandler{}
}
