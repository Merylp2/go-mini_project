package controllers

import (
	"go_movie-ticket/lib/database"
	"go_movie-ticket/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreatePriceController(c echo.Context) error {
	price := models.Price{}
	c.Bind(&price)

	Price, err := database.CreatePrice(price)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success create Price",
		"Data":    Price,
	})
}

func UpdatePriceByIdController(c echo.Context) error {
	PriceId := c.Param("id")

	price := models.Price{}
	c.Bind(&price)

	price, err := database.UpdatePrice(price, PriceId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success update price",
		"Data":    price,
	})
}

func DeletePriceByIdController(c echo.Context) error {
	PriceId := c.Param("id")

	_, err := database.DeletePrice(PriceId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success delete price",
	})
}
