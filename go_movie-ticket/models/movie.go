package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Movie struct {
	gorm.Model
	MovieId     int       `json:"movie_id" form:"movie_id"`
	Title       string    `json:"title" form:"title"`
	Genre       string    `json:"genre" form:"genre"`
	Description string    `json:"description" form:"description"`
	Rating      float64   `json:"rating" form:"rating"`
	ShowTimes   time.Time `json:"show_times" form:"show_times" gorm:"many2many:movie_show_times"`
	
}
