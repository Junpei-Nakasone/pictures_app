package data

import (
	"pictures_app/api/api008/domain"
	"pictures_app/environment/db"
)

// AddNewUser ユーザー新規登録API
func AddNewUser(param domain.RequestParam) (domain.ResponseParam, error) {

	db := db.CreateDBConnection()
	defer db.Close()

	data := domain.User{
		UserName:     param.UserName,
		UserPassword: param.Password,
		EmailAddress: param.EmailAddress,
		Note:         param.Note,
		// IconImageは未実装
		IconImage: nil,
	}

	err := db.Table("users").
		Create(&data).Error

	if err != nil {
		return domain.ResponseParam{}, err
	}

	res := domain.ResponseParam{
		UserID:    data.UserID,
		UserName:  data.UserName,
		Note:      data.Note,
		IconImage: data.IconImage,
	}

	return res, nil
}
