package util

import (
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/steinfletcher/apitest"
)

func MethodTest(t *testing.T, apiEndpoint string, app *echo.Echo, allowedMethod string) {
	if allowedMethod != "Get" {
		t.Run("Getメソッドは許可されていない", func(t *testing.T) {
			apitest.New().
				Handler(app).
				Get(apiEndpoint).
				Expect(t).
				Status(http.StatusMethodNotAllowed).
				End()
		})
	}

	if allowedMethod != "Post" {
		t.Run("Postメソッドは許可されていない", func(t *testing.T) {
			apitest.New().
				Handler(app).
				Post(apiEndpoint).
				Expect(t).
				Status(http.StatusMethodNotAllowed).
				End()
		})
	}

	if allowedMethod != "Put" {
		t.Run("Putメソッドは許可されていない", func(t *testing.T) {
			apitest.New().
				Handler(app).
				Put(apiEndpoint).
				Expect(t).
				Status(http.StatusMethodNotAllowed).
				End()
		})
	}

	if allowedMethod != "Delete" {
		t.Run("Deleteメソッドは許可されていない", func(t *testing.T) {
			apitest.New().
				Handler(app).
				Delete(apiEndpoint).
				Expect(t).
				Status(http.StatusMethodNotAllowed).
				End()
		})
	}
}
