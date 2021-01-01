package data

import (
	"pictures_app/api/api006/domain"
	"pictures_app/api/api006/usecase/repository"
	"time"

	"github.com/jinzhu/gorm"
)

type serviceRepository struct {
	db *gorm.DB
}

// NewServiceRepository 初期化処理
func NewServiceRepository(db *gorm.DB) repository.ServiceRepository {
	return &serviceRepository{db: db}
}

func (t *serviceRepository) AddImage(userID int, imageURL string) error {
	data := domain.Pictures{
		UserID:      userID,
		ImageURL:    imageURL,
		PublishedAt: time.Now(),
	}

	err := t.db.Table("pictures").
		Create(&data).Error

	return err
}
