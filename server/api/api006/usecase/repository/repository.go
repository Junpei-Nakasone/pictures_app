package repository

import "pictures_app/api/api006/domain"

type ServiceRepository interface {
	AddImage(domain.Picture) error
}
