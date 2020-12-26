package infra

import (
	"pictures_app/api/api002/domain"
	"pictures_app/environment/db"
)

// FetchImages 画像を投稿順に取得する
func FetchImages() ([]domain.Picture, error) {
	db := db.CreateDBConnection()
	defer db.Close()

	result := []domain.Picture{}

	err := db.Order("published_at DESC").Find(&result).Error

	return result, err
}
