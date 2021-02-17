package repository

import "pictures_app/api/api010/domain"

// ServiceRepository インターフェース
type ServiceRepository interface {
	FetchImagesByPictureID(param domain.RequestParam) (domain.Picture, error)
}
