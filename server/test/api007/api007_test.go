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

func TestApi007Test001(t *testing.T) {

	environment.SetEnvVariables()

	db := db.CreateDBConnection()
	app := environment.NewApp(db)

	prepareDB(db.DB())

	util.MethodTest(t, "/login", app.App, "Post")

	apitest.New().
		Handler(app.App).
		Post("/login").
		Header("Content-Type", "application/json").
		BodyFromFile("testdata/request/test001.golden").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestApi007Test002(t *testing.T) {

	environment.SetEnvVariables()

	db := db.CreateDBConnection()
	app := environment.NewApp(db)

	prepareDB(db.DB())

	util.MethodTest(t, "/login", app.App, "Post")

	apitest.New().
		Handler(app.App).
		Post("/login").
		Header("Content-Type", "application/json").
		BodyFromFile("testdata/request/test002.golden").
		Expect(t).
		Status(http.StatusNotFound).
		End()
}
