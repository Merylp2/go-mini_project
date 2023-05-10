package models

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title       string    `json:"title" form:"title"`
	Genre       string    `json:"genre" form:"genre"`
	Duration    string    `json:"duration" form:"duration"`
	Director    string    `json:"director" form:"director"`
	Description string    `json:"description" form:"description"`
	Price       int       `json:"price" form:"price"`
	Seat        int       `json:"seat" form:"seat"`
	Rating      int       `json:"rating" form:"rating"`
	ShowTimes   time.Time `json:"show_times" form:"show_times"`
}
