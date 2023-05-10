package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string  `json:"name" form:"name"`
	NoHP     string  `json:"no_hp" form:"no_hp"`
	Username string  `json:"username" form:"username"`
	Password string  `json:"-" form:"password"`
	Saldo    uint     `json:"saldo" form:"saldo"`
	Token    string  `json:"-" form:"token"`
	Order    []Order `json:"-" gorm:"foreignKey:UserID"`
}

type Token struct {
	Username string `json:"username" form:"username"`
	Token    string `json:"token" form:"token"`
}
