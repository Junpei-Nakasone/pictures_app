package repository

import "pictures_app/api/api005/domain"

// ServiceRepository インターフェース
type ServiceRepository interface {
	FetchViewCategories() ([]domain.ResponseParam, error)
}
