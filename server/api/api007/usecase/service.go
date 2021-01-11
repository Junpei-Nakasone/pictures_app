package usecase

import (
	"pictures_app/api/api007/domain"
	"pictures_app/api/api007/usecase/repository"
	"pictures_app/api/common"
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
	userData, err := s.rep.FetchUserData(param)
	if err != nil {
		return domain.ResponseParam{}, err
	}

	err = common.VerifyPassword(userData.UserPassword, param.Password)
	if err != nil {
		return domain.ResponseParam{}, err
	}

	res := domain.ResponseParam{
		UserID:    userData.UserID,
		UserName:  userData.UserName,
		Note:      userData.Note,
		IconImage: userData.IconImage,
	}
	return res, nil
}
