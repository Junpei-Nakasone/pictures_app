package util

import (
	"database/sql"
)

// PrepareDB データベースにデータを投入する
func PrepareDB(db *sql.DB) {
	// db := db.CreateDBConnection()

	// fixtures, err := testfixtures.New(
	// 	testfixtures.Database(db.DB()),
	// 	testfixtures.Dialect("mysql"),
	// 	testfixtures.Directory("testdata/fixtures"),
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if err = fixtures.Load(); err != nil {
	// 	log.Fatal(err)
	// }
}
