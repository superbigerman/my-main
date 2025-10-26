package service

type DataService struct {
	producer  interface{ Produce() ([]string, error) }
	presenter interface{ Present(data []string) }
}

func NewDataService(producer interface{ Produce() ([]string, error) },
	presenter interface{ Present(data []string) }) *DataService {
	return &DataService{
		producer:  producer,
		presenter: presenter,
	}
}

func (s *DataService) Process() error {
	data, err := s.producer.Produce()
	if err != nil {
		return err
	}
	s.presenter.Present(data)
	return nil

}
