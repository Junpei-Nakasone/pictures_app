package service

import (
	"nuxt-dadjokes/api/api002/domain"
	"nuxt-dadjokes/api/api002/infra"
)

// FetchImages 画像を投稿順に取得する
func FetchImages() (result []domain.Picture, err error) {
	result, err = infra.FetchImages()
	return result, err
}
