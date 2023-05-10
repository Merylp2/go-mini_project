package database

import (
	"go_movie-ticket/config"
	"go_movie-ticket/models"
)

func GetMovie() ([]models.Movie, error) {
	movie := []models.Movie{}

	if err := config.DB.Find(&movie).Error; err != nil {
		return movie, err
	}
	return movie, nil
}

func GetMovieById(id any) (models.Movie, error) {
	var movie models.Movie

	if err := config.DB.Where("id = ?", id).First(&movie).Error; err != nil {
		return movie, err
	}

	return movie, nil
}

func UpdateMovie(id uint) (models.Movie, error) {
	var movies models.Movie
	if err := config.DB.Where("id = ?", id).Save(&movies).Error; err != nil {
		return movies, err
	}

	return movies, nil
}
