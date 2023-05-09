package database

import (
	"go_movie-ticket/config"
	"go_movie-ticket/middlewares"
	"go_movie-ticket/models"
)

func LoginUsers(user *models.User) (interface{}, error) {
	var err error
	if err = config.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(user).Error; err != nil {
		return nil, err
	}

	var token models.Token
	token.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}

	return token, nil
}

// Create user
func CreateUser(user models.User) (models.User, error) {
	err := config.DB.Create(&user).Error

	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// Get all user
func GetUser() (users []models.User, err error) {
	err = config.DB.Find(&users).Error

	if err != nil {
		return []models.User{}, err
	}
	return
}

// Get user by id
func GetUserById(id any) (models.User, error) {
	var user models.User

	err := config.DB.Where("id = ?", id).First(&user).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// Update user by id
func UpdateUser(user models.User, id any) (models.User, error) {
	err := config.DB.Table("users").Where("id = ?", id).Updates(&user).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// Delete user by id
func DeleteUser(id any) (interface{}, error) {
	err := config.DB.Where("id = ?", id).Delete(&models.User{}).Error

	if err != nil {
		return nil, err
	}

	return "success delete user", nil
}
