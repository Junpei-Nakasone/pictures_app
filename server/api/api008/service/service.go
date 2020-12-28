package service

import (
	"pictures_app/api/api008/domain"
	"pictures_app/api/api008/infra/data"
)

// AddNewUser ユーザー新規登録API
func AddNewUser(param domain.RequestParam) (domain.ResponseParam, error) {
	res, err := data.AddNewUser(param)
	if err != nil {
		return domain.ResponseParam{}, err
	}

	return res, nil
}
