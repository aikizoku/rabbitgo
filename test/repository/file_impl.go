package repository

import (
	"fmt"
	"os"
)

type file struct {
	rootPath string
}

func (r *file) CreateDir(path string) {
	if err := os.MkdirAll(r.rootPath+path, 0755); err != nil {
		panic(err)
	}
}

func (r *file) WriteFile(path string, body string) {
	file, err := os.OpenFile(r.rootPath+path, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, body)
	if err != nil {
		panic(err)
	}
}

func (r *file) RemoveAll() {
	if err := os.RemoveAll(r.rootPath); err != nil {
		panic(err)
	}
}

// NewFile ... Fileを作成する
func NewFile(rootPath string) File {
	return &file{
		rootPath: rootPath,
	}
}
