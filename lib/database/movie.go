package database

import (
	"go_movie-ticket/config"
	"go_movie-ticket/models"
)

func CreateMovie(movie models.Movie) (models.Movie, error) {
	err := config.DB.Create(&movie).Error

	if err != nil {
		return models.Movie{}, err
	}
	return movie, nil
}

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

func UpdateMovie(id int) (models.Movie, error) {
	var movies models.Movie
	if err := config.DB.Where("id = ?", id).Updates(&movies).Error; err != nil {
		return movies, err
	}

	return movies, nil
}

func DeleteMovie(id any) (interface{}, error) {
	err := config.DB.Where("id = ?", id).Delete(&models.Movie{}).Error

	if err != nil {
		return nil, err
	}

	return "success delete movie", nil
}
