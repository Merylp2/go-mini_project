package usecase

import (
	"go_movie-ticket/lib/database"
	"go_movie-ticket/models"
)

func GetAllMovie() (movie []models.Movie, err error) {
	movie, err = database.GetMovie()
	if err != nil {
		return movie, err
	}

	return movie, nil
}

func GetMovieById(id int) (models.Movie, error) {
	movie, err := database.GetMovieById(id)
	if err != nil {
		return movie, err
	}

	return movie, nil
}
