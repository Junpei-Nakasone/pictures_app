package data

import (
	"pictures_app/api/api002/domain"
	"pictures_app/api/api002/usecase/repository"

	"github.com/jinzhu/gorm"
	// "pictures_app/environment/db"
)

type serviceRepository struct {
	db *gorm.DB
}

// NewServiceRepository 初期化処理
func NewServiceRepository(db *gorm.DB) repository.ServiceRepository {
	return &serviceRepository{db: db}
}

// FetchImages 画像を投稿順に取得する
func (t *serviceRepository) FetchImages() ([]domain.Picture, error) {
	// db := db.CreateDBConnection()
	// defer db.Close()

	result := []domain.Picture{}

	err := t.db.Order("published_at DESC").Find(&result).Error

	return result, err
}
