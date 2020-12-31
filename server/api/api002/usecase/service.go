package usecase

import (
	"pictures_app/api/api002/domain"
	"pictures_app/api/api002/usecase/repository"
)

type service struct {
	rep repository.ServiceRepository
}

type Service interface {
	FetchImages() ([]domain.Picture, error)
}

func NewService(r repository.ServiceRepository) Service {
	return &service{rep: r}
}

// FetchImages 画像を投稿順に取得する
func (s *service) FetchImages() (result []domain.Picture, err error) {
	result, err = s.rep.FetchImages()
	return result, err
}
