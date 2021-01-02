package repository

import "pictures_app/api/api004/domain"

// ServiceRepository インターフェース
type ServiceRepository interface {
	FetchPrefectureCategories() ([]domain.ResponseParam, error)
}
