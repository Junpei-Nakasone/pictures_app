package data

import (
	"pictures_app/api/api006/domain"
	"pictures_app/api/api006/usecase/repository"

	"github.com/jinzhu/gorm"
)

type serviceRepository struct {
	db *gorm.DB
}

// NewServiceRepository 初期化処理
func NewServiceRepository(db *gorm.DB) repository.ServiceRepository {
	return &serviceRepository{db: db}
}

func (t *serviceRepository) AddImage(pictureData domain.Picture) error {

	err := t.db.Table("pictures").
		Create(&pictureData).Error

	return err
}
