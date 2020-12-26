package main

import "nuxt-dadjokes/environment/router"

func main() {
	e := router.NewRouter()

	e.Start(":9999")
}
