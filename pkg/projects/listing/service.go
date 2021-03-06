package listing

type Repository interface {
	ListAll() []Project
	GetById(id int) Project
}

type Service struct {
	r Repository
}

func (s *Service) ListAll() []Project {
	return s.r.ListAll()
}

func (s *Service) GetById(id int) Project {
	return s.r.GetById(id)
}

func NewService(r Repository) *Service {
	return &Service{r}
}
