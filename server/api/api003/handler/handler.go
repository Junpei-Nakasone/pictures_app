package handler

import (
	"net/http"
	"pictures_app/api/api003/domain"
	"pictures_app/environment/db"

	"github.com/labstack/echo"
)

// AddJoke is for new joke.
func AddJoke(c echo.Context) error {

	db := db.CreateDBConnection()
	defer db.Close()

	param := domain.RequestParam{}
	if err := c.Bind(&param); err != nil {
		return err
	}

	err := db.Table("jokes").Create(&param).Error
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, "New Joke added.")
}
