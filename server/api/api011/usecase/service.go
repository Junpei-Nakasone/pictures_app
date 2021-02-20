package usecase

import (
	"pictures_app/api/api011/domain"
	"pictures_app/api/api011/usecase/repository"
)

type service struct {
	rep repository.ServiceRepository
}

type Service interface {
	FetchUserData(param domain.RequestParam) (domain.UserData, error)
}

func NewService(r repository.ServiceRepository) Service {
	return &service{rep: r}
}

// FetchImages 画像を投稿順に取得する
func (s *service) FetchUserData(param domain.RequestParam) (result domain.UserData, err error) {
	userData, err := s.rep.FetchUserData(param)
	if err != nil {
		return domain.UserData{}, err
	}

	pictureData, err := s.rep.FetchPicturesByUserID(*param.UserID)

	userData.PostedPictures = pictureData

	return userData, err
}
