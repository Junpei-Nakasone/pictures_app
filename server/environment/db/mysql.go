package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// CreateDBConnection 環境変数で指定したDBと接続する
func CreateDBConnection() *gorm.DB {

	// ENVIRONMENTがlocalの場合はローカルのMySQLコンテナと接続する
	if os.Getenv("ENVIRONMENT") == "local" {
		db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/test_database?parseTime=true")
		if err != nil {
			fmt.Println("error occured at db connection")
			return nil
		}
		db.LogMode(true)

		return db
	}

	// ENVIRONTMENTがproductionの場合は本番環境のDBと接続する
	if os.Getenv("ENVIRONMENT") == "production" {

		DBMS := os.Getenv("DBMS")
		USER := os.Getenv("USER")
		PASSWORD := os.Getenv("PASSWORD")
		HOST := os.Getenv("HOST")
		DBNAME := os.Getenv("DBNAME")
		CONNECT := USER + ":" + PASSWORD + "@" + HOST + "/" + DBNAME + "?parseTime=true"
		db, err := gorm.Open(DBMS, CONNECT)

		if err != nil {
			fmt.Println("error occured at db connection")
			return nil
		}

		return db
	}

	fmt.Println("環境変数を設定してください")
	return nil
}
