package data

import (
	"pictures_app/api/api011/domain"
	"pictures_app/api/api011/usecase/repository"

	"github.com/jinzhu/gorm"
)

type serviceRepository struct {
	db *gorm.DB
}

// NewServiceRepository 初期化処理
func NewServiceRepository(db *gorm.DB) repository.ServiceRepository {
	return &serviceRepository{db: db}
}

// FetchUserData ユーザー情報を取得する
func (t *serviceRepository) FetchUserData(param domain.RequestParam) (domain.UserData, error) {

	result := domain.UserData{}

	err := t.db.Table(`users`).
		Where(`user_id = ?`, param.UserID).
		Find(&result).Error

	return result, err
}

func (t *serviceRepository) FetchPicturesByUserID(userID int) ([]domain.Picture, error) {
	result := []domain.Picture{}

	err := t.db.Table(`pictures`).
		Where(`user_id = ?`, userID).
		Find(&result).Error

	return result, err
}
