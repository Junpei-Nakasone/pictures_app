package data

import (
	"pictures_app/api/api012/domain"
	"pictures_app/api/api012/usecase/repository"

	"github.com/jinzhu/gorm"
)

type serviceRepository struct {
	db *gorm.DB
}

// NewServiceRepository 初期化処理
func NewServiceRepository(db *gorm.DB) repository.ServiceRepository {
	return &serviceRepository{db: db}
}

// DeletePicture ユーザー情報を取得する
func (t *serviceRepository) DeletePicture(param domain.RequestParam) error {

	data := domain.Picture{}
	t.db.Table(`pictures`).
		Where(`picture_id = ?`, param.PictureID).
		Find(&data)

	err := t.db.Where("picture_id = ?", param.PictureID).Delete(&data).Error

	return err
}
