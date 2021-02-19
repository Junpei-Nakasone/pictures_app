package api007

import (
	"database/sql"
	"log"
	"net/http"
	"pictures_app/environment"
	"pictures_app/environment/db"
	"pictures_app/test/util"
	"testing"

	// "pictures_app/environment/router"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/steinfletcher/apitest"
)

func prepareDB(db *sql.DB) {
	// db := db.CreateDBConnection()

	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
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

func TestApi007Test(t *testing.T) {

	// prepareDB()
	db := db.CreateDBConnection()
	app := environment.NewApp(db)

	prepareDB(db.DB())

	util.MethodTest(t, "/login", app.App, "Post")

	apitest.New().
		Handler(app.App).
		Post("/login").
		Header("Content-Type", "application/json").
		BodyFromFile("testdata/test001.golden").
		Expect(t).
		Status(http.StatusOK).
		End()
}
