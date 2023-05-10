package config

import (
	"fmt"
	"go_movie-ticket/models"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {

	config := Config{
		DB_Username: "alta",
		DB_Password: "root",
		DB_Port:     "3306",
		DB_Host:     "192.168.1.5",
		DB_Name:     "go_movie-ticket",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	InitialMigration()

}

func InitialMigration() {
	DB.AutoMigrate(&models.User{}, &models.Movie{}, &models.Order{})
}
