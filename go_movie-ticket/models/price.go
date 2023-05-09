package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Price struct {
	gorm.Model
	Price     float64		`json:"price" form:"price"`
	Day		  time.Weekday	`json:"day" form:"day"`
}