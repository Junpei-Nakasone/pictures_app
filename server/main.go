package main

import (
	"pictures_app/env"
	"pictures_app/environment"
	"pictures_app/environment/db"
)

func main() {
	env.SetEnvVariables()
	db := db.CreateDBConnection()
	defer db.Close()

	app := environment.NewApp(db)

	app.App.Start(":4567")
}
