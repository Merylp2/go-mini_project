package models

import (
	"gorm.io/gorm"
)

type Price struct {
	gorm.Model
	MovieId int     `json:"movie_id" form:"movie_id"`
	Price   float64 `json:"price" form:"price"`
}