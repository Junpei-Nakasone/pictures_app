package api007

import (
	"net/http"
	"nuxt-dadjokes/environment/router"
	"testing"

	"github.com/steinfletcher/apitest"
)

func TestApi007Test(t *testing.T) {
	apitest.New().
		Handler(router.NewRouter()).
		Post("/login").
		Header("Content-Type", "application/json").
		BodyFromFile("testdata/test001.golden").
		Expect(t).
		Status(http.StatusOK).
		End()
}
