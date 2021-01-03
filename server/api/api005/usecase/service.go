package usecase

import (
	"pictures_app/api/api005/domain"
	"pictures_app/api/api005/usecase/repository"
)

type service struct {
	rep repository.ServiceRepository
}

type Service interface {
	FetchViewCategories() ([]domain.ResponseParam, error)
}

func NewService(r repository.ServiceRepository) Service {
	return &service{rep: r}
}

// FetchImages 画像を投稿順に取得する
func (s *service) FetchViewCategories() (result []domain.ResponseParam, err error) {
	result, err = s.rep.FetchViewCategories()
	return result, err
}
