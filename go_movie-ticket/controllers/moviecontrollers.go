package controllers

import (
	"go_movie-ticket/lib/database"
	"go_movie-ticket/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateMovieController(c echo.Context) error {
	movie := models.Movie{}
	c.Bind(&movie)

	movie, err := database.CreateMovie(movie)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success create movie",
		"Data":    movie,
	})
}

func GetMoviesControllers(c echo.Context) error {
	movies, err := database.GetMovie()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, movies)
}

func GetMovieByIdController(c echo.Context) error {
	MovieId := c.Param("id")

	movie, err := database.GetMovieById(MovieId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success get Movie",
		"Data":    movie,
	})
}

func UpdateMovieByIdController(c echo.Context) error {
	MovieId := c.Param("id")

	movie := models.Movie{}
	c.Bind(&movie)

	movie, err := database.UpdateMovie(movie, MovieId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success update movie",
		"Data":    movie,
	})
}

func DeleteMovieByIdController(c echo.Context) error {
	movieId := c.Param("id")

	_, err := database.DeleteMovie(movieId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success delete movie",
	})
}
