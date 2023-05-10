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
		Quantity: req.Quantity,
	}

	if movie.Seat == 0 {
		return models.Order{}, echo.NewHTTPError(400, "Tiket sudah habis")
	}

	order, _ := database.CheckOrder(regisReq)
	if order.UserID >= 10 {
		return models.Order{}, echo.NewHTTPError(400, "Mencapai batas pembelian tiket")
	}

	if int(user.Saldo) < movie.Price {
		return models.Order{}, echo.NewHTTPError(400, "Saldo tidak cukup")
	} else {
		user.Saldo -= movie.Price

		err = database.UpdateUsers(&user)
		if err != nil {
			return models.Order{}, err
		}
	}

	order, err = database.CreateOrder(regisReq)
	if err != nil {
		return models.Order{}, err
	}

	movie.Seat -= req.Quantity

	_, err = database.UpdateMovie(int(regisReq.MovieID))
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
