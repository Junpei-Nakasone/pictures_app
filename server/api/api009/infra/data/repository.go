package data

import (
	"pictures_app/api/api009/domain"
	"pictures_app/api/api009/usecase/repository"

	"github.com/jinzhu/gorm"
)

type serviceRepository struct {
	db *gorm.DB
}

// NewServiceRepository 初期化処理
func NewServiceRepository(db *gorm.DB) repository.ServiceRepository {
	return &serviceRepository{db: db}
}

// FetchImages 画像を投稿順に取得する
func (t *serviceRepository) FetchImagesByViewCategoryCd(param domain.RequestParam) ([]domain.Picture, error) {

	result := []domain.Picture{}

	err := t.db.Where(`view_category_cd = ?`, param.ViewCategoryCd).Find(&result).Error

	return result, err
}
