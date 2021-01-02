package repository

import "pictures_app/api/api003/domain"

// ServiceRepository インターフェース
type ServiceRepository interface {
	FetchImagesByPrefectureCd(param domain.RequestParam) ([]domain.Picture, error)
}
