package controllers

import (
	"fmt"
	"go_movie-ticket/lib/database"
	"go_movie-ticket/models/payload"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetMoviesControllers(c echo.Context) error {
	movies, err := database.GetMovie()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "success get all movies",
		Data:    movies,
	})
}

func GetMovieByIdController(c echo.Context) error {
	Id := c.Param("id")

	movie, err := database.GetMovieById(Id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: fmt.Sprintf("success get movie  %s", movie.Title),
		Data:    movie,
	})
}
