package repository

import "pictures_app/api/api009/domain"

// ServiceRepository インターフェース
type ServiceRepository interface {
	FetchImagesByViewCategoryCd(param domain.RequestParam) ([]domain.Picture, error)
}
