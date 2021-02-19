package main

import "pictures_app/api/common"

// アプリケーションとは別で動く文字列をハッシュ化するための関数
func main() {

	str := "password1"
	hash, _ := common.HashPassword(str)

	println(hash)
}
