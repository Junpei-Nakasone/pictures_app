package api008

import (
	"log"
	"pictures_app/environment/db"

	// "pictures_app/environment/router"

	"github.com/go-testfixtures/testfixtures/v3"
)

func prepareDB() {
	db := db.CreateDBConnection()

	fixtures, err := testfixtures.New(
		testfixtures.Database(db.DB()),
		testfixtures.Dialect("mysql"),
		testfixtures.Directory("testdata/fixtures"),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err = fixtures.Load(); err != nil {
		log.Fatal(err)
	}
}

// func TestApi008Test(t *testing.T) {

// 	prepareDB()

// 	apitest.New().
// 		Handler(router.NewRouter()).
// 		Post("/addNewUser").
// 		Header("Content-Type", "application/json").
// 		BodyFromFile("testdata/test001.golden").
// 		Expect(t).
// 		Status(http.StatusOK).
// 		End()

// }
