package repository

import "pictures_app/api/api002/domain"

// ServiceRepository インターフェース
type ServiceRepository interface {
	FetchImages() ([]domain.Picture, error)
}
