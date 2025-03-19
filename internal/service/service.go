package service

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) GetRoomCleanlinessStatus(image string) (string, error) {
	return "clean", nil
}
