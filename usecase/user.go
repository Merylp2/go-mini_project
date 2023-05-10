package usecase

import (
	"go_movie-ticket/lib/database"
	"go_movie-ticket/models"
	"go_movie-ticket/models/payload"
)

func GetUser(id int) (user models.User, err error) {
	user, err = database.GetUserById(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func UpdateUser(id int, req payload.UpdateUser) (user models.User, err error) {
	user, err = database.GetUserById(id)
	if err != nil {
		return user, err
	}

	userRq := models.User{
		Name:     req.Name,
		NoHP:     req.NoHP,
		Username: req.Username,
		Password: req.Password,
	}

	user, err = database.UpdateUser(id, userRq)
	if err != nil {
		return user, err
	}

	return user, nil
}

func DeleteUser(id int) (user models.User, err error) {

	user, err = database.DeleteUser(id)
	if err != nil {
		return user, err
	}

	return user, nil
}
