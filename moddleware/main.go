package main

import (
	"moddleware/database"
	"moddleware/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8000")

	return
}
