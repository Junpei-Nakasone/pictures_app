package usecase

import (
	"pictures_app/api/api004/domain"
	"pictures_app/api/api004/usecase/repository"
)

type service struct {
	rep repository.ServiceRepository
}

type Service interface {
	FetchPrefectureCategories() ([]domain.ResponseParam, error)
}

func NewService(r repository.ServiceRepository) Service {
	return &service{rep: r}
}

// FetchImages 画像を投稿順に取得する
func (s *service) FetchPrefectureCategories() (result []domain.ResponseParam, err error) {
	result, err = s.rep.FetchPrefectureCategories()
	return result, err
}
