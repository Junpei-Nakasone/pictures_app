package service

import (
	"pictures_app/api/api002/domain"
	"pictures_app/api/api002/infra"
)

// FetchImages 画像を投稿順に取得する
func FetchImages() (result []domain.Picture, err error) {
	result, err = infra.FetchImages()
	return result, err
}
