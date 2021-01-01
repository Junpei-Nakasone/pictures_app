package data

import (
	"pictures_app/api/api008/domain"
	"pictures_app/api/api008/usecase/repository"

	"github.com/jinzhu/gorm"
)

type serviceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) repository.ServiceRepository {
	return &serviceRepository{db: db}
}

// AddNewUser ユーザー新規登録API
func (t *serviceRepository) AddNewUser(param domain.RequestParam) (domain.ResponseParam, error) {

	data := domain.User{
		UserName:     param.UserName,
		UserPassword: param.Password,
		EmailAddress: param.EmailAddress,
		Note:         param.Note,
		// IconImageは未実装
		IconImage: nil,
	}

	err := t.db.Table("users").
		Create(&data).Error

	if err != nil {
		return domain.ResponseParam{}, err
	}

	res := domain.ResponseParam{
		UserID:    data.UserID,
		UserName:  data.UserName,
		Note:      data.Note,
		IconImage: data.IconImage,
	}

	return res, nil
}
