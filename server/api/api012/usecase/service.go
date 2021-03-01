package usecase

import (
	"pictures_app/api/api012/domain"
	"pictures_app/api/api012/usecase/repository"
)

type service struct {
	rep repository.ServiceRepository
}

type Service interface {
	DeletePicture(param domain.RequestParam) error
}

func NewService(r repository.ServiceRepository) Service {
	return &service{rep: r}
}

// FetchImages 画像を投稿順に取得する
func (s *service) DeletePicture(param domain.RequestParam) error {
	return s.rep.DeletePicture(param)
}
