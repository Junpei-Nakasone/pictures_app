package repository

import "pictures_app/api/api012/domain"

// ServiceRepository インターフェース
type ServiceRepository interface {
	DeletePicture(param domain.RequestParam) error
}
