package main

import (
	"test/22_Middleware/Praktikum/config"
	"test/22_Middleware/Praktikum/middlewares"
	"test/22_Middleware/Praktikum/routes"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// create a new echo instance
	config.InitDB()
	e := routes.New()

	middlewares.LogMiddlewares(e)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))

}
