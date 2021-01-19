package usecase

import (
	"pictures_app/api/api006/domain"
	"pictures_app/api/api006/usecase/repository"
)

type service struct {
	rep repository.ServiceRepository
}

type Service interface {
	AddImage(domain.Picture) error
}

func NewService(r repository.ServiceRepository) Service {
	return &service{rep: r}
}

func (s *service) AddImage(pictureData domain.Picture) error {
	return s.rep.AddImage(pictureData)
}
