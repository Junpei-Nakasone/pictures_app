package api002

import (
	"database/sql"
	"log"
	"net/http"
	"pictures_app/environment"
	"pictures_app/environment/db"
	"pictures_app/test/util"
	"testing"

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

func TestApi002Test001(t *testing.T) {

	environment.SetEnvVariables()

	db := db.CreateDBConnection()
	app := environment.NewApp(db)

	prepareDB(db.DB())

	util.MethodTest(t, "/fetchLatestImages", app.App, "Get")

	apitest.New().
		Handler(app.App).
		Get("/fetchLatestImages").
		Header("Content-Type", "application/json").
		BodyFromFile("testdata/request/test001.golden").
		Expect(t).
		BodyFromFile("testdata/response/response001.golden").
		Status(http.StatusOK).
		End()
}

func TestApi002Test002(t *testing.T) {

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
