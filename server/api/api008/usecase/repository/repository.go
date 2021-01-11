package repository

import "pictures_app/api/api008/domain"

type ServiceRepository interface {
	AddNewUser(param domain.RequestParam, hashedPassword string) (domain.ResponseParam, error)
}
