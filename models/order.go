package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	Name      string    `json:"name" form:"name"`
	Title     string    `json:"title" form:"title"`
	ShowTimes time.Time `json:"show_times" form:"show_times" gorm:"many2many:movie_show_times"`
	Quantity  int       `json:"quantity" form:"quantity"`
	Kursi     string    `json:"kursi" form:"kursi"`
	Price     float64   `json:"price" form:"price"`
	UserID    uint      `json:"user_id"`
	MovieID   uint      `json:"movie_id"`
}
