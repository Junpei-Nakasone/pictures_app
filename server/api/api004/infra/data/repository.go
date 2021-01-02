package data

import (
	"pictures_app/api/api004/domain"
	"pictures_app/api/api004/usecase/repository"

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
func (t *serviceRepository) FetchPrefectureCategories() ([]domain.ResponseParam, error) {

	result := []domain.ResponseParam{}

	err := t.db.Table(`prefecture_categories`).
		Order(`sort_no ASC`).Find(&result).Error

	return result, err
}
