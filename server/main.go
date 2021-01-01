package main

import (
	"pictures_app/environment"
	"pictures_app/environment/db"
)

func main() {
	db := db.CreateDBConnection()
	defer db.Close()

	app := environment.NewApp(db)

	app.App.Start(":9999")
}
