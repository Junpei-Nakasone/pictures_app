package usecase

import "pictures_app/api/api006/usecase/repository"

type service struct {
	rep repository.ServiceRepository
}

type Service interface {
	AddImage(userID int, imageURL string) error
}

func NewService(r repository.ServiceRepository) Service {
	return &service{rep: r}
}

func (s *service) AddImage(userID int, imageURL string) error {
	return s.rep.AddImage(userID, imageURL)
}
