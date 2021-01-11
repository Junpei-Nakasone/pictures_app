package repository

import "pictures_app/api/api007/domain"

type ServiceRepository interface {
	FetchUserData(domain.RequestParam) (domain.UserData, error)
}
