package env

import "os"

// SetEnvVariables 環境変数を設定
func SetEnvVariables() {
	os.Setenv("DBMS", "mysqlest")
	os.Setenv("USER", "rootFROMSETENV")
	os.Setenv("PASSWORD", "root")
	os.Setenv("PROTOCOL", "tcp(localhost:3306)")
	os.Setenv("DBNAME", "test_database")
	os.Setenv("testparameter", "thisistestparameterfronSETENV")
}
