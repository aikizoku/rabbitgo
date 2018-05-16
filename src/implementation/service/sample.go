package implementation

import "github.com/aikizoku/go-web-template/src/repository"

type SampleService struct {
	r *repository.Sample
}

func NewSampleService() *SampleService {
	return &SampleService{}
}

func (s *SampleService) Hoge() {

}
