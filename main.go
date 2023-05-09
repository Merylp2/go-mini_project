package main

import (
	"go_movie-ticket/config"
	m "go_movie-ticket/middlewares"
	"go_movie-ticket/routes"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	config.InitDB()
	e := routes.New()
	m.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8080"))

}
