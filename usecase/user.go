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
		return models.User{}, err
	}

	if req.Name != "" {
		user.Name = req.Name
	}

	if req.NoHP != "" {
		user.NoHP = req.NoHP
	}

	if req.Username != "" {
		user.Username = req.Username
	}

	if req.Password != "" {
		user.Password = req.Password
	}

	err = database.UpdateUser(&user)
	if err != nil {
		return models.User{}, err
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

func TopupUser(id, saldo uint) (models.User, error) {
	user, err := database.GetUserById(id)
	if err != nil {
		return models.User{}, err
	}

	user.Saldo += saldo
	err = database.UpdateUser(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, err
}
