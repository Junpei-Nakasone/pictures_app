package service

import (
	"pictures_app/api/api007/domain"
	"pictures_app/api/api007/infra/data"
)

// FetchUserData ログインするユーザー情報を取得
func FetchUserData(param domain.RequestParam) (domain.ResponseParam, error) {
	return data.FetchUserData(param)
}
