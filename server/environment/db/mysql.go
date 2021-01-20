package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// CreateDBConnection returns db
func CreateDBConnection() *gorm.DB {
	DBMS := "mysql"
	// DBMS := os.Getenv("DBMS")
	// USER := "root"
	// USER := os.Getenv("USER")
	// PASSWORD := "root"
	// PASSWORD := os.Getenv("PASSWORD")
	USER := "admin"
	PASSWORD := "password"
	// ローカルのGoからDockerのMySQLに繋ぐ時はコンテナ名ではなくlocalhostで指定する
	// PROTOCOL := "tcp(localhost:3306)"
	// PROTOCOL := os.Getenv("PROTOCOL")
	// PROTOCOL := "tcp(mysql:3306)"
	// RDS
	PROTOCOL := "tcp(rds-pictures-0106.ci3srwjbei5c.ap-northeast-1.rds.amazonaws.com:3306)"
	DBNAME := "test_database"
	// DBNAME := os.Getenv("DBNAME")
	CONNECT := USER + ":" + PASSWORD + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)
	// db, err := gorm.Open("mysql", "admin:password@tcp(database-1.ci3srwjbei5c.ap-northeast-1.rds.amazonaws.com:3306)/test_database?parseTime=true")
	if err != nil {
		fmt.Println("error occured at db connection")
		return nil
	}
	db.LogMode(true)

	return db
}
