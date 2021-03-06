package data

import (
	"pictures_app/api/api010/domain"
	"pictures_app/api/api010/usecase/repository"

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
func (t *serviceRepository) FetchImagesByPictureID(param domain.RequestParam) (domain.Picture, error) {

	result := domain.Picture{}

	err := t.db.Table(`pictures`).
		Select(`pictures.picture_id
				,		pictures.user_id
				,		pictures.image_url
				,		pictures.published_at
				,		users.user_name
				,		users.icon_image
				,		pictures.image_note
			`).
		Joins(`INNER JOIN users ON
		 pictures.user_id = users.user_id`).
		Where(`picture_id = ?`, param.PictureID).
		Find(&result).Error

	return result, err
}
