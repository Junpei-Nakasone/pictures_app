package repository

import "pictures_app/api/api011/domain"

// ServiceRepository インターフェース
type ServiceRepository interface {
	FetchUserData(param domain.RequestParam) (domain.UserData, error)
	FetchPicturesByUserID(userID int) ([]domain.Picture, error)
}
