package controllers

import (
	"go_movie-ticket/lib/database"
	"go_movie-ticket/middlewares"
	"go_movie-ticket/models/payload"
	"go_movie-ticket/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateOrderController(c echo.Context) error {
	req := payload.OrderRequest{}

	c.Bind(&req)

	id := middlewares.ExtractTokenUserId(c)

	order, err := usecase.CreateOrder(id, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success create order",
		"Data":    order,
	})
}

func GetOrderByUserIdControllers(c echo.Context) error {

	id := middlewares.ExtractTokenUserId(c)

	orders, err := database.GetOrderByUserId(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "success get order",
		Data:    orders,
	})
}
