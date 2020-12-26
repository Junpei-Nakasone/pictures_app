package handler

import (
	"net/http"
	"nuxt-dadjokes/api/api005/domain"
	"nuxt-dadjokes/environment/db"

	"github.com/labstack/echo"
)

// DeleteJoke is for new joke.
func DeleteJoke(c echo.Context) error {

	db := db.CreateDBConnection()
	defer db.Close()

	param := domain.RequestParam{}
	if err := c.Bind(&param); err != nil {
		return err
	}

	err := db.Table("jokes").Where("id = ?", param.ID).Delete(&param).Error

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, "That Joke deleted.")
}
