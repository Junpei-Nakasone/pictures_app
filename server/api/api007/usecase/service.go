package usecase

import (
	"pictures_app/api/api007/domain"
	"pictures_app/api/api007/usecase/repository"
)

type service struct {
	rep repository.ServiceRepository
}

type Service interface {
	FetchUserData(param domain.RequestParam) (domain.ResponseParam, error)
}

func NewService(r repository.ServiceRepository) Service {
	return &service{rep: r}
}

// FetchUserData ログインするユーザー情報を取得
func (s *service) FetchUserData(param domain.RequestParam) (domain.ResponseParam, error) {
	return s.rep.FetchUserData(param)
}
