package main

import (
	"workshoptdd/config"
	"workshoptdd/routes"
)

func main() {
	db := config.InitDatabase("host=localhost user=postgres password=postgres dbname=workshop_tdd")
	app := routes.InitRoutes(db)

	app.Run(":8000")
}
