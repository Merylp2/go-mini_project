package usecase

import (
	"go_movie-ticket/lib/database"
	"go_movie-ticket/models"
	"go_movie-ticket/models/payload"

	"github.com/labstack/echo"
)

func CreateOrder(id int, req payload.OrderRequest) (models.Order, error) {

	user, err := database.GetUserById(id)
	if err != nil {
		return models.Order{}, err
	}

	movie, err := database.GetMovieById(req.MovieID)
	if err != nil {
		return models.Order{}, err
	}

	regisReq := models.Order{
		UserID:   user.ID,
		MovieID:  uint(req.MovieID),
		Quantity: uint(req.Quantity),
	}

	if movie.Seat == 0 {
		return models.Order{}, echo.NewHTTPError(400, "Tiket sudah habis")
	}

	if uint(user.Saldo) < (movie.Price * req.Quantity) {
		return models.Order{}, echo.NewHTTPError(400, "Saldo tidak cukup")
	} else {
		user.Saldo -= (movie.Price * req.Quantity)

		err = database.UpdateUsers(&user)
		if err != nil {
			return models.Order{}, err
		}
	}

	_, err = database.CreateOrder(regisReq)
	if err != nil {
		return models.Order{}, err
	}

	movie.Seat -= req.Quantity

	_, err = database.UpdateMovie(uint(regisReq.MovieID))
	if err != nil {
		return models.Order{}, err
	}

	order, err := database.GetOrderByUserID(int(user.ID))
	if err != nil {
		return models.Order{}, err
	}

	return order, nil
}

func GetOrderByUserId(id int) (order []models.Order, err error) {
	user, err := database.GetUserById(id)
	if err != nil {
		return nil, err
	}

	order, err = database.GetOrderByUserId(int(user.ID))
	if err != nil {
		return nil, err
	}

	return order, nil
}
