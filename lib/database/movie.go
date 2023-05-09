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

func GetMovie() (movie models.Movie, err error) {
	err = config.DB.First(&movie).Error
	if err != nil {
		return movie, err
	}
	return movie, nil
}

func GetMovieById(id any) (models.MovieDetail, error) {
	var movie models.Movie
	var movieD models.MovieDetail

	err := config.DB.Where("id = ?", id).First(&movie).Error

	if err != nil {
		return models.MovieDetail{}, err
	}

	movieD.MovieId = movie.MovieId
	movieD.Title = movie.Title
	movieD.Genre = movie.Genre
	movieD.Director = movie.Director
	movieD.Description = movie.Description
	movieD.Duration = movie.Duration
	movieD.Rating = movie.Rating
	movieD.ShowTimes = movie.ShowTimes

	return movieD, nil
}

func UpdateMovie(movie models.Movie, id any) (models.Movie, error) {
	var existingMovie models.Movie
	err := config.DB.Table("movies").Where("id = ?", id).First(&existingMovie).Error
	if err != nil {
		return models.Movie{}, err
	}

	err = config.DB.Table("movies").Model(&existingMovie).Updates(&movie).Error
	if err != nil {
		return models.Movie{}, err
	}

	return existingMovie, nil
}

func DeleteMovie(id any) (interface{}, error) {
	err := config.DB.Where("id = ?", id).Delete(&models.Movie{}).Error

	if err != nil {
		return nil, err
	}

	return "success delete movie", nil
}
