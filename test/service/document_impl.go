package service

import (
	"fmt"
	"strings"

	"github.com/aikizoku/skgo/test/model"
	"github.com/aikizoku/skgo/test/repository"
)

type document struct {
	fRepo repository.File
	tRepo repository.TemplateClient
}

func (s *document) RemoveAll() {
	s.fRepo.RemoveAll()
}

func (s *document) Distributes(tmplPath string, apis []*model.API) {
	for _, api := range apis {
		s.distribute(tmplPath, api)
	}
}

func (s *document) distribute(tmplPath string, api *model.API) {
	str := s.tRepo.GetMarged("template/api.tmpl", api)

	dirPath := fmt.Sprintf("/%s", strings.Replace(api.Overview.URI, "/", "_", -1))
	s.fRepo.CreateDir(dirPath)

	s.fRepo.WriteFile(fmt.Sprintf("%s/%s.md", dirPath, api.Name), str)
}

// NewDocument ... Documentを作成する
func NewDocument(fRepo repository.File, tRepo repository.TemplateClient) Document {
	return &document{
		fRepo: fRepo,
		tRepo: tRepo,
	}
}
