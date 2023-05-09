package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Movie struct {
	gorm.Model
	MovieId		int			`json:"movie_id" form:"movie_id"`
	Title		string		`json:"title" form:"title"`
	Genre		string		`json:"genre" form:"genre"`
	Duration	string		`json:"duration" form:"duration"`
	Director	string		`json:"-" form:"director"`
	Description	string		`json:"-" form:"description"`
	Rating      float64		`json:"-" form:"rating"`
	ShowTimes   time.Time	`json:"-" form:"show_times"`
}

type MovieDetail struct {
	MovieId     int			`json:"movie_id" form:"movie_id"`
	Title       string		`json:"title" form:"title"`
	Genre       string		`json:"genre" form:"genre"`
	Director    string		`json:"director" form:"director"`
	Description string		`json:"description" form:"description"`
	Duration    string		`json:"duration" form:"duration"`
	Rating      float64		`json:"rating" form:"rating"`
	ShowTimes   time.Time	`json:"show_times" form:"show_times"`
}

func (MovieDetail) TableName() string {
	return "movies"
}


