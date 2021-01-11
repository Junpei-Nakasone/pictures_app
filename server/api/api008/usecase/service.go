package usecase

import (
	"pictures_app/api/api008/domain"
	"pictures_app/api/api008/usecase/repository"
	"pictures_app/api/common"
)

type service struct {
	rep repository.ServiceRepository
}

type Service interface {
	AddNewUser(param domain.RequestParam) (domain.ResponseParam, error)
}

func NewService(r repository.ServiceRepository) Service {
	return &service{rep: r}
}

// AddNewUser ユーザー新規登録API
func (s *service) AddNewUser(param domain.RequestParam) (domain.ResponseParam, error) {
	hashedPassword, err := common.HashPassword(*param.Password)
	if err != nil {
		return domain.ResponseParam{}, err
	}
	res, err := s.rep.AddNewUser(param, hashedPassword)
	if err != nil {
		return domain.ResponseParam{}, err
	}

	return res, nil
}
