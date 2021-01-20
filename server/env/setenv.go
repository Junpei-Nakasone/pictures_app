package env

import "os"

// SetEnvVariables 環境変数を設定
func SetEnvVariables() {
	os.Setenv("DBMS", "mysql")
	// os.Setenv("USER", "root")
	os.Setenv("USER", "admin")
	// os.Setenv("PASSWORD", "root")
	os.Setenv("PASSWORD", "password")
	// os.Setenv("PROTOCOL", "tcp(localhost:3306)")
	os.Setenv("PROTOCOL", "rds-pictures-0106.ci3srwjbei5c.ap-northeast-1.rds.amazonaws.com")
	os.Setenv("DBNAME", "test_database")
}
