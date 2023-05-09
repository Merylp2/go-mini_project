package controllers

import (
	"go_movie-ticket/lib/database"
	"go_movie-ticket/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateOrderController(c echo.Context) error {
	order := models.Order{}
	c.Bind(&order)

	order, err := database.CreateOrder(order)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success create order",
		"Data":    order,
	})
}

func GetOrderControllers(c echo.Context) error {
	orders, err := database.GetOrder()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, orders)
}
