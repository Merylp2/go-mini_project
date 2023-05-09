package models

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name" validate:"required"`
	NoHP     string `json:"no_hp" form:"no_hp" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type Token struct {
	Token string `json:"token" form:"token"`
}

type CustomValidator struct {
	Validators *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.Validators.Struct(i)

	if err != nil {
		var sb strings.Builder
		sb.WriteString("Validation error:\n")

		for _, err := range err.(validator.ValidationErrors) {
			sb.WriteString(fmt.Sprintf("- %s\n", err))
		}

		return echo.NewHTTPError(http.StatusBadRequest, sb.String())
	}

	return nil
}
