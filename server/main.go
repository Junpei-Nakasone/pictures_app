package main

import (
	"pictures_app/environment/router"
)

func main() {
	e := router.NewRouter()

	e.Start(":9999")
}
