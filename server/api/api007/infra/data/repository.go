package data

import (
	"pictures_app/api/api007/domain"
	"pictures_app/environment/db"
)

// FetchUserData ログインするユーザー情報を取得
func FetchUserData(param domain.RequestParam) (domain.ResponseParam, error) {

	result := domain.ResponseParam{}

	db := db.CreateDBConnection()
	defer db.Close()

	err := db.Table("users").
		Select(`
			user_id
		, user_name
		, note
		, icon_image`).
		Where(`user_name = ?
		AND user_password = ?`, param.UserName, param.Password).
		Find(&result).Error

	return result, err

}
