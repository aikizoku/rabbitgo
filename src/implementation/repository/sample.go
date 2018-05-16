package implementation

type SampleRepository struct {
}

func NewSampleRepository() *SampleRepository {
	return &SampleRepository{}
}

func (s *SampleRepository) hoge() {
}
