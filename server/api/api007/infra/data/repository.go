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
func (t *serviceRepository) FetchUserData(param domain.RequestParam) (domain.UserData, error) {

	result := domain.UserData{}

	err := t.db.Table("users").
		Select(`
			user_id
		, user_name
		, user_password
		, note
		, icon_image`).
		Where(`email_address = ?`, param.EmailAddress).
		Find(&result).Error

	return result, err

}
