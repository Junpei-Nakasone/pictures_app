package data

import (
	"pictures_app/api/api007/domain"
	"pictures_app/api/api007/usecase/repository"

	"github.com/jinzhu/gorm"
)

type serviceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) repository.ServiceRepository {
	return &serviceRepository{db: db}
}

// FetchUserData ログインするユーザー情報を取得
func (t *serviceRepository) FetchUserData(param domain.RequestParam) (domain.ResponseParam, error) {

	result := domain.ResponseParam{}

	err := t.db.Table("users").
		Select(`
			user_id
		, user_name
		, note
		, icon_image`).
		Where(`user_name = ?
		AND user_password = ?`, param.UserName, param.Password).
		Find(&result).Error

	return result, err

}
