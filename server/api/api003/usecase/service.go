package usecase

import (
	"pictures_app/api/api003/domain"
	"pictures_app/api/api003/usecase/repository"
)

type service struct {
	rep repository.ServiceRepository
}

type Service interface {
	FetchImagesByPrefectureCd(param domain.RequestParam) ([]domain.Picture, error)
}

func NewService(r repository.ServiceRepository) Service {
	return &service{rep: r}
}

// FetchImages 画像を投稿順に取得する
func (s *service) FetchImagesByPrefectureCd(param domain.RequestParam) (result []domain.Picture, err error) {
	result, err = s.rep.FetchImagesByPrefectureCd(param)
	return result, err
}
